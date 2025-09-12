package module

type Module struct {
	Number int
	Name   string
	Params map[string]string
}

func newModule(number int, name string, params map[string]string) *Module {
	mod := Module{
		Number: number,
		Name:   name,
		Params: params,
	}

	return &mod
}
