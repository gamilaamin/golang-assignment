package config

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // The address of the server
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // port
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // Advanced configuration
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // The database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // The database user name
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // Database password
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // The maximum number of connections in free
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // Open the maximum number of connections to the database
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
