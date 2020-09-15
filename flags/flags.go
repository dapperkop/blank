package flags

import (
	"flag"

	"github.com/dapperkop/blank/consts"
	"github.com/dapperkop/blank/database/migrations/create"
	"github.com/dapperkop/blank/database/migrations/downto"
	"github.com/dapperkop/blank/database/migrations/migrate"
	"github.com/dapperkop/blank/database/migrations/upto"
	"github.com/dapperkop/blank/publish/config"
	"github.com/dapperkop/blank/publish/logs"
	"github.com/dapperkop/blank/publish/migrations"
)

// SetupAndParse func ...
func SetupAndParse(mode *string) {
	// Init migrations flags
	flag.Var(create.GetFlag(), "create", "[-create=name]\t\t\t\t\t\tCreates new migration file with the current timestamp.")
	flag.Var(downto.GetFlag(), "down-to", "[-down-to=version]\t\t\t\t\tRoll back to a specific version.")
	flag.Var(migrate.GetFlag(), "migrate", migrate.GetFlagUsage())
	flag.Var(upto.GetFlag(), "up-to", "[-up-to=version]\t\t\t\t\tMigrate the DB to a specific version.")

	// Init publish flags
	flag.Var(config.GetFlag(), "publish-config", "[-publish-config=dir]")
	flag.Var(logs.GetFlag(), "publish-logs", "[-publish-logs=dir]")
	flag.Var(migrations.GetFlag(), "publish-migrations", "[-publish-migrations=dir]")

	// Init app flag
	flag.StringVar(mode, "mode", consts.DefaultAppMode, "[-mode=env]\t\t\t\t\t\tEnvironment name. {{env}}.toml file must exist.")

	// Parse all flags from the command-line
	flag.Parse()
}
