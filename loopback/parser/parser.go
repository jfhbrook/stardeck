package parser

import (
	"bytes"
	"regexp"
)

func ParseModuleOutput(output []byte) (*Module, error) {
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
