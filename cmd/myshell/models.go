package main

type CommandType string

const (
	INTERNAL CommandType = "INTERNAL"
	SYSTEM   CommandType = "SYSTEM"
)

type Command struct {
	Type    CommandType
	Command string
	F       func([]string) error
	Man     string
	Path    string
}

type WrongArgumentsError struct {
	msg string
}

func (e *WrongArgumentsError) Error() string { return e.msg }

func CreateSystemCommand(command string, path string) Command {
	return Command{
		Type:    SYSTEM,
		Command: command,
		Path:    path,
		F: func(args []string) error {
			return nil
		},
	}
}
