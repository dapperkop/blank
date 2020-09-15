package create

import (
	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	"github.com/pressly/goose"
)

var createFlag types.NullCreate

// GetFlag func ...
func GetFlag() *types.NullCreate {
	return &createFlag
}

// Need func ...
func Need() bool {
	return createFlag.Valid
}

// Run func ...
func Run(migrationsDir string) {
	if !Need() {
		return
	}

	var (
		dir = filesystem.Getwd() + "/" + migrationsDir
		err error
	)

	err = goose.Create(database.DB, dir, createFlag.Name, "go")

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}
