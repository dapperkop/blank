package types

import (
	"fmt"

	"github.com/dapperkop/blank/consts"
)

// Dirs type ...
type Dirs struct {
	ConfigDir     string `toml:"config"`
	LogsDir       string `toml:"logs"`
	MigrationsDir string `toml:"migrations"`
}

// NullConfig type ...
type NullConfig struct {
	Dir   string
	Valid bool
}

// String func ...
func (config *NullConfig) String() string {
	return fmt.Sprint(*config)
}

// Set func ...
func (config *NullConfig) Set(value string) error {
	if value == "" {
		config.Dir = consts.DefaultPublishConfigDir
	} else {
		config.Dir = value
	}

	config.Valid = true

	return nil
}

// NullLogs type ...
type NullLogs struct {
	Dir   string
	Valid bool
}

// String func ...
func (logs *NullLogs) String() string {
	return fmt.Sprint(*logs)
}

// Set func ...
func (logs *NullLogs) Set(value string) error {
	if value == "" {
		logs.Dir = consts.DefaultPublishLogsDir
	} else {
		logs.Dir = value
	}

	logs.Valid = true

	return nil
}

// NullMigrations type ...
type NullMigrations struct {
	Dir   string
	Valid bool
}

// String func ...
func (migrations *NullMigrations) String() string {
	return fmt.Sprint(*migrations)
}

// Set func ...
func (migrations *NullMigrations) Set(value string) error {
	if value == "" {
		migrations.Dir = consts.DefaultPublishMigrationsDir
	} else {
		migrations.Dir = value
	}

	migrations.Valid = true

	return nil
}
