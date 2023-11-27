package config

import "time"

var TimeLocation *time.Location

type DbConn struct {
	Open int `koanf:"open"`
	TTL  int `koanf:"ttl"`
	Idle int `koanf:"idle"`
}

type DB struct {
	Host       string `koanf:"host"`
	Port       int    `koanf:"port"`
	Username   string `koanf:"user"`
	Password   string `koanf:"pass"`
	Name       string `koanf:"name"`
	Connection DbConn `koanf:"conn"`
}

type App struct {
	Name     string `koanf:"name"`
	Url      string `koanf:"url"`
	Port     int    `koanf:"port"`
	Env      string `koanf:"env"`
	Debug    bool   `koanf:"debug"`
	Timezone string `koanf:"tz"`
}

type Api struct {
	Key string `koanf:"key"`
}

type Cfg struct {
	App App `koanf:"app"`
	Db  DB  `koanf:"db"`
	Api Api `koanf:"api"`
}

var Config *Cfg
