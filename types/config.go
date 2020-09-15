package types

// APIServer type ...
type APIServer struct {
	HTTP HTTP `toml:"http"`
}

// Config type ...
type Config struct {
	Database  Database  `toml:"database"`
	Logger    Logger    `toml:"logger"`
	APIServer APIServer `toml:"api-server"`
}

// Database type ...
type Database struct {
	Driver   string `toml:"driver"`
	Hostname string `toml:"hostname"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	DBPort   int    `toml:"dbport"`
}

// HTTP type ...
type HTTP struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// Logger type ...
type Logger struct {
	Debug    bool   `toml:"debug"`
	Timezone string `toml:"timezone"`
}
