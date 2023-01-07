package cmd

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
	}
}
