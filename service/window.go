package service

func makeSetWindowNameCommand(name string) *Command {
	cmd := Command{
		Type:  SetWindowNameCommand,
		Value: name,
	}

	return &cmd
}
