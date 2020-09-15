package helpers

import (
	"github.com/dapperkop/blank/consts"
	"github.com/dapperkop/blank/types"
)

// InitConfig func ...
func InitConfig() types.Config {
	return types.Config{
		Database: types.Database{
			Driver:   consts.DefaultDatabaseDriver,
			Hostname: consts.DefaultDatabaseHostname,
			Username: consts.DefaultDatabaseUsername,
			Password: consts.DefaultDatabasePassword,
			DBName:   consts.DefaultDatabaseDBName,
			DBPort:   consts.DefaultDatabaseDBPort,
		},
		Logger: types.Logger{
			Debug:    consts.DefaultLoggerDebug,
			Timezone: consts.DefaultLoggerTimezone,
		},
		APIServer: types.APIServer{
			HTTP: types.HTTP{
				Host: consts.DefaultAPIHTTPHost,
				Port: consts.DefaultAPIHTTPPort,
			},
		},
	}
}

// InitDirs func ...
func InitDirs() types.Dirs {
	return types.Dirs{
		ConfigDir:     consts.DefaultPublishConfigDir,
		LogsDir:       consts.DefaultPublishLogsDir,
		MigrationsDir: consts.DefaultPublishMigrationsDir,
	}
}
