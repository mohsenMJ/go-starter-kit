package cmd

func Execute(args []string) {
	command := args[1]
	switch command {
	case "migrate":
		Migrate()
	case "seed":
		RunSeeders()
	}
}
