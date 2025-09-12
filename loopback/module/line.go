package module

import (
	"regexp"
	"strconv"
)

type lineParser struct {
	line           []byte
	lineNo         int
	columnNo       int
	spacesRe       *regexp.Regexp
	eqRe           *regexp.Regexp
	braceRe        *regexp.Regexp
	moduleNumberRe *regexp.Regexp
	identifierRe   *regexp.Regexp
	valueRe        *regexp.Regexp
}

func newLineParser(line []byte, lineNo int) *lineParser {
	parser := lineParser{
		line:           line,
		lineNo:         lineNo,
		columnNo:       0,
		spacesRe:       regexp.MustCompile(`^\s+`),
		eqRe:           regexp.MustCompile(`^=`),
		braceRe:        regexp.MustCompile(`^\{`),
		moduleNumberRe: regexp.MustCompile(`^\d+`),
		identifierRe:   regexp.MustCompile(`^[a-zA-Z\-_]+`),
		valueRe:        regexp.MustCompile(`^[a-zA-Z\-_\d]+`),
	}

	return &parser
}

func (p *lineParser) match(re *regexp.Regexp) []int {
	if p.columnNo >= len(p.line) {
		return nil
	}
	loc := re.FindIndex(p.line[p.columnNo:])
	if loc == nil {
		return nil
	}
	loc[0] = loc[0] + p.columnNo
	loc[1] = loc[1] + p.columnNo
	p.columnNo = loc[1]
	return loc
}

func (p *lineParser) expect(re *regexp.Regexp, message string) ([]int, *ParseError) {
	loc := p.match(re)
	if loc == nil {
		return nil, p.parseError(CodeExpect, message)
	}
	return loc, nil
}

func (p *lineParser) substring(loc []int) string {
	return string(p.line[loc[0]:loc[1]])
}

func (p *lineParser) parseError(code string, message string) *ParseError {
	err := ParseError{
		Code:     code,
		Message:  message,
		LineNo:   p.lineNo,
		ColumnNo: p.columnNo,
	}

	return &err
}

func (p *lineParser) module() (*Module, *ParseError) {
	moduleNo, err := p.moduleNumber()
	if err != nil {
		return nil, err
	}

	p.spaces()

	name, err := p.moduleName()

	if err != nil {
		return nil, err
	}

	params := make(map[string]string)

	for {
		err := p.spaces()
		if err != nil {
			return newModule(moduleNo, name, params), nil
		}

		if p.match(p.braceRe) != nil {
			return newModule(moduleNo, name, params), p.parseError(CodeComplex, "Encountered complex params")
		}

		key, value, err := p.param()
		if err != nil {
			if err.Code == CodeNone {
				return newModule(moduleNo, name, params), nil
			}
			return nil, err
		}

		params[key] = value
	}
}

func (p *lineParser) spaces() *ParseError {
	_, err := p.expect(p.spacesRe, "Expected whitespace")
	return err
}

func (p *lineParser) eq() *ParseError {
	_, err := p.expect(p.eqRe, "Expected '='")
	return err
}

func (p *lineParser) value() (string, *ParseError) {
	loc, err := p.expect(p.valueRe, "Expected value")

	if err != nil {
		return "", err
	}

	return p.substring(loc), nil
}

func (p *lineParser) moduleNumber() (int, *ParseError) {
	loc, err := p.expect(p.moduleNumberRe, "Expected module number")
	if err != nil {
		return -1, err
	}
	no, err := p.number(loc)
	if err != nil {
		return -1, err
	}

	return no, nil
}

func (p *lineParser) number(loc []int) (int, *ParseError) {
	no, err := strconv.Atoi(p.substring(loc))
	if err != nil {
		return -1, p.parseError(CodeNumber, err.Error())
	}
	return no, nil
}

func (p *lineParser) moduleName() (string, *ParseError) {
	loc, err := p.expect(p.identifierRe, "Expected module name")
	if err != nil {
		return "", err
	}

	return p.substring(loc), nil
}

func (p *lineParser) paramName() (string, *ParseError) {
	loc, err := p.expect(p.identifierRe, "Expected parameter name")
	if err != nil {
		return "", err
	}

	return p.substring(loc), nil
}

func (p *lineParser) param() (string, string, *ParseError) {
	key, err := p.paramName()

	if err != nil {
		return "", "", p.parseError(CodeNone, err.Message)
	}

	err = p.eq()

	if err != nil {
		return "", "", err
	}

	value, err := p.value()

	if err != nil {
		return "", "", err
	}

	return key, value, nil
}
