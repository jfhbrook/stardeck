package parser

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
