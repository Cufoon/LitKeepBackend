package conf

import "fmt"

type Conf struct {
	Server  `yaml:"server"`
	LogZap  `yaml:"log_zap"`
	LogFile `yaml:"log_file"`
	Redis   `yaml:"redis"`
	Mysql   `yaml:"db"`
	Gorm    `yaml:"gorm"`
	Mode    string `yaml:"mode"`
}

type Server struct {
	Port string `yaml:"port"`
}

type LogZap struct {
	Level  int `yaml:"level"`
	Output int `yaml:"output"`
}

type LogFile struct {
	Filename   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

type Redis struct {
	Enable   string `yaml:"enable"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

type Mysql struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db_name"`
	Parameters string `yaml:"parameters"`
}

func (d *Mysql) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.Parameters,
	)
}

type Gorm struct {
	Debug        string `yaml:"debug"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	MaxOpenConns int    `yaml:"max_open_connections"`
	MaxIdleConns int    `yaml:"max_idle_connections"`
}
