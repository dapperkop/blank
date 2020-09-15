package migrations

import (
	"strings"

	"github.com/dapperkop/blank/database/migrations/create"
	"github.com/dapperkop/blank/database/migrations/downto"
	"github.com/dapperkop/blank/database/migrations/migrate"
	"github.com/dapperkop/blank/database/migrations/upto"
	"github.com/dapperkop/blank/logger"
	"github.com/dapperkop/blank/publish"
	"github.com/dapperkop/blank/types"
	"github.com/pressly/goose"
)

var migrationsDir = publish.LoadDirs().MigrationsDir

// Need func ...
func Need() bool {
	return create.Need() || downto.Need() || migrate.Need() || upto.Need()
}

// Run func ...
func Run() {
	if !Need() {
		return
	}

	create.Run(migrationsDir)
	downto.Run(migrationsDir)
	migrate.Run(migrationsDir)
	upto.Run(migrationsDir)
}

// Setup func ...
func Setup(config types.Database) {
	var err error

	goose.SetLogger(logger.Logger)
	err = goose.SetDialect(config.Driver)

	if err != nil {
		logger.Logger.Fatalln(err)
	}
}

// Status func ...
func Status() string {
	var status []string

	if create.Need() {
		status = append(status, "create")
	}

	if downto.Need() {
		status = append(status, "downto")
	}

	if migrate.Need() {
		status = append(status, "migrate")
	}

	if upto.Need() {
		status = append(status, "upto")
	}

	return strings.Join(status, " & ")
}
