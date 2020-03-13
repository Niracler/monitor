package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode       string
	ImageSavePath string

	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	DBType        string
	DBHost        string
	DBName        string
	DBUser        string
	DBPassword    string
	DBTablePrefix string

	RedisHost string
	RedisPass string

	ResourceType string
)

func init() {
	var err error
	Cfg, err = ini.Load("setting/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'setting/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustString("8000")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	ImageSavePath = sec.Key("IMAGE_SAVE_PATH").MustString("media/images/")
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	ResourceType = sec.Key("RESOURCE_TYPE").MustString("")
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	DBType = sec.Key("DB_TYPE").String()
	DBName = sec.Key("DB_NAME").String()
	DBUser = sec.Key("DB_USER").String()
	DBPassword = sec.Key("DB_PASSWORD").String()
	DBHost = sec.Key("DB_HOST").String()
	DBTablePrefix = sec.Key("DB_TABLE_PREFIX").String()

	RedisHost = sec.Key("REDIS_HOST").MustString("127.0.0.1:6379")
	RedisPass = sec.Key("REDIS_PASS").MustString("123456")
}
