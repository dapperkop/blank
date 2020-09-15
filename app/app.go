package app

import (
	"os"

	"github.com/dapperkop/blank/apiserver"
	"github.com/dapperkop/blank/config"
	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/database/migrations"
	"github.com/dapperkop/blank/flags"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/publish"
	"github.com/dapperkop/blank/types"
)

var (
	// Config var ...
	Config types.Config

	// Mode var ...
	Mode string
)

func getStatus() string {
	var status string

	switch {
	case publish.Need():
		status = "Publish [" + publish.Status() + "] launched"
	case migrations.Need():
		status = "Migration [" + migrations.Status() + "] launched"
	default:
		status = "API Server configured"
	}

	status += " in \"" + Mode + "\" mode"

	return status
}

// Run func ...
func Run() {
	// Setup app
	setup()

	// Print app status info
	logger.Logger.Infoln(getStatus())

	// Run
	switch {
	case publish.Need():
		// Run publish
		publish.Run(Mode)

		os.Exit(0)
	case migrations.Need():
		// Run migrations
		migrations.Run()

		os.Exit(0)
	}
}

func setup() {
	// Setup and parse flags
	flags.SetupAndParse(&Mode)

	// Get config by mode
	Config = config.Get(Mode)

	// Setup logger
	logger.Setup(Config.Logger)
	// Setup database
	database.Setup(Config.Database)

	// Setup
	switch {
	case publish.Need():
		// Setup publish
		publish.Setup()
	case migrations.Need():
		// Setup migrations
		migrations.Setup(Config.Database)
	default:
		// Setup apiserver
		apiserver.Setup(Config.APIServer)
	}
}
