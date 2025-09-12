package module

import (
	"bytes"
	"regexp"
)

func Parse(output []byte) (*Module, error) {
	closingBraceRe := regexp.MustCompile(`^\s+}\s*$`)
	lines := bytes.Split(output, []byte("\n"))

	if len(lines) == 0 {
		return nil, notFoundError(0)
	}

	i := 0
	line := lines[0]

	advance := func() *ParseError {
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

		if module != nil {
			if module.Name == "module-loopback" {
				return module, nil
			}
		}

		if err != nil {
			if err.Code == CodeComplex {
				for !closingBraceRe.Match(line) {
					if err := advance(); err != nil {
						return nil, err
					}
				}
			} else {
				if err := advance(); err != nil {
					return nil, err
				}
			}
		} else {
			if err := advance(); err != nil {
				return nil, err
			}
		}
	}
}
