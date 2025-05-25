package builtin

type Command interface {
	GetCmdName() string
	Handle(args []string)
}

func GetBuiltInMap() map[string]Command {
	builtinMap := map[string]Command{
		"echo": &Echo{},
		"exit": &Exit{},
		"type": &Type{},
		"pwd":  &Pwd{},
	}

	return builtinMap
}
