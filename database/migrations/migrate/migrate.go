package migrate

import (
	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/helpers/filesystem"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/types"
	"github.com/pressly/goose"
)

var migrateFlag types.NullMigrate

// GetFlag func ...
func GetFlag() *types.NullMigrate {
	return &migrateFlag
}

// GetFlagUsage func ...
func GetFlagUsage() string {
	var (
		joinedMigrateValues = types.MigrateValues.Join("|")
		usage               = "[-migrate=" + joinedMigrateValues + "]\tRun the database migrations."
	)

	return usage
}

// Need func ...
func Need() bool {
	return migrateFlag.Valid
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

	switch migrateFlag.Migrate {
	case "up":
		err = goose.Up(database.DB, dir)
	case "up-by-one":
		err = goose.UpByOne(database.DB, dir)
	case "down":
		err = goose.Down(database.DB, dir)
	case "redo":
		err = goose.Redo(database.DB, dir)
	case "reset":
		err = goose.Reset(database.DB, dir)
	case "status":
		err = goose.Status(database.DB, dir)
	case "version":
		err = goose.Version(database.DB, dir)
	}

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}
