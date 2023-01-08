package cmd

var Commands []Command

type Command struct {
	Run         string
	Description string
	Example     string
	Author      string
}

func Execute(args []string) {
	// Get Command Name
	command := args[1]

	switch command {
	case "test":
		CmdTestHander()
	case "migrate":
		CmdMigrateHandler()
	case "seed":
		CmdRunSeedersHandler()
	case "make:model":
		CmdMakeModelHandler(args)
	case "help":
		CmdHelpHandler()
	}
}
