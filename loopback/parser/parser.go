package parser

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

const (
	codeExpect string = "EXPECT"
	codeNumber = "NUMBER"
	codeComplex = "COMPLEX"
	codeNotFound = "NOT_FOUND"
)

type ParseError struct {
	code string
	message string
	lineNo int
	columnNo int
}

func (err ParseError) Error() string {
	return fmt.Sprintf("Parse error at %d:%d: %s", err.lineNo, err.columnNo, err.message)
}

type Module struct {
	number int
	name string
	params map[string]string
}

func newModule(number int, name string, params map[string]string) *Module {
	mod := Module {
		number: number,
		name: name,
		params: params,
	}

	return &mod
}

type lineParser struct {
	line []byte
	lineNo int
	columnNo int
	spacesRe *regexp.Regexp
	eqRe *regexp.Regexp
	braceRe *regexp.Regexp
	moduleNumberRe *regexp.Regexp
	identifierRe *regexp.Regexp
	valueRe *regexp.Regexp
}

func newLineParser(line []byte, lineNo int) *lineParser {
	parser := lineParser {
		line: line,
		lineNo: lineNo,
		columnNo: 0,
		spacesRe: regexp.MustCompile(`^ +`),
		eqRe: regexp.MustCompile(`^=`),
		braceRe: regexp.MustCompile(`^\{`),
		moduleNumberRe: regexp.MustCompile(`^\d+`),
		identifierRe: regexp.MustCompile(`^[a-zA-Z\-_]+`),
		valueRe: regexp.MustCompile(`^[a-zA-Z\-_\d]+`),
	}

	return &parser
}

func (p *lineParser) match(re *regexp.Regexp) []int {
	loc := re.FindIndex(p.line[p.columnNo:])
	if loc == nil {
		return nil
	}
	loc[0] += p.columnNo
	loc[1] += p.columnNo
	p.columnNo += loc[1]
	return loc
}

func (p *lineParser) expect(re *regexp.Regexp, message string) ([]int, *ParseError) {
	loc := p.match(re)
	if loc == nil {
		return nil, p.parseError(codeExpect, message)
	}
	return loc, nil
}

func (p *lineParser) substring(loc []int) string {
	return string(p.line[loc[0]:loc[1]])
}

func (p *lineParser) parseError(code string, message string) *ParseError {
	err := ParseError {
		code: code,
		message: message,
		lineNo: p.lineNo,
		columnNo: p.columnNo,
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
			return newModule(moduleNo, name, params), p.parseError(codeComplex, "Encountered complex params")
		}

		key, value, err := p.param()
		if err != nil {
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
		return -1, p.parseError(codeNumber, err.Error())
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
		return "", "", err
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


func notFoundError(lineNo int) *ParseError {
	err := ParseError {
		code: codeNotFound,
		message: "Module not found",
		lineNo: lineNo,
		columnNo: 0,
	}

	return &err
}


func parse(output []byte) (*Module, error) {
	closingBraceRe := regexp.MustCompile(`^ +\}+ *$`)
	lines := bytes.Split(output, []byte("\n"))

	if len(lines) == 0 {
		return nil, notFoundError(0)
	}

	i := 0
	line := lines[0]

	advance := func() error {
		i += 1
		if i >= len(lines) {
			return notFoundError(i)
		}
		line = lines[i]
		return nil
	}

	for {
		parser := newLineParser(line, i + 1)
		module, err := parser.module()

		if module != nil && module.name == "module-loopback" {
			return module, err
		}

		if err := advance(); err != nil {
			return nil, err
		}

		if err != nil && err.code == codeComplex {
				for !closingBraceRe.Match(line) {
					if err := advance(); err != nil {
						return nil, err
					}
				}
		}
	}
}
