package config

import (
	"encoding/json"
	"os"
	"wms/log"
)

// HttpConfig todo: Http
type HttpConfig struct {
	Listen string `json:"listen"`
	//Secret string `json:"secret"`
}

// MysqlConfig todo: Mysql
type MysqlConfig struct {
	Database string `json:"database"`
	Settings string `json:"settings"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Idle     int    `json:"idle"`
	Max      int    `json:"max"`
	ShowSQL  bool   `json:"showSQL"`
}

// RedisConfig todo: Redis
type RedisConfig struct {
	Address  string `json:"address"`  // 访问地址 <ip/domain>:port
	Username string `json:"username"` // 用户名 TBD
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库
}

// EmailConfig todo: Email
type EmailConfig struct {
	SMTP     string `json:"smtp"`
	PORT     int    `json:"port"`
	NUM      string `json:"num"`
	PASSWORD string `json:"password"`
}

// Config todo: 配置项
type Config struct {
	Env       string      `json:"env"`
	Http      HttpConfig  `json:"http"`
	Mysql     MysqlConfig `json:"mysql"`
	Redis     RedisConfig `json:"redis"`
	Email     EmailConfig `json:"email"`
	PublicKey string      `json:"public_key"`
}

// Instance todo: 默认配置实例
var Instance *Config

// ENV环境变量
const (
	//Delos todo: 生产环境
	Delos = "Delos"
	//Prod todo: 测试环境
	Prod = "Prod"
	//Dev todo: 开发环境
	Dev = "Dev" // 开发环境
)

// NewConfig 从文件中加载一个配置实例
func NewConfig(file string) (*Config, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// Init 初始化加载配置信息
func Init(file string) error {
	conf, err := NewConfig(file)
	if err != nil {
		log.Error("NewConfig error: ", err.Error())
		return err
	}

	// 如果没有配置env，默认是开发环境
	if len(conf.Env) == 0 {
		conf.Env = Dev
	}

	if len(conf.Http.Listen) == 0 {
		conf.Http.Listen = "127.0.0.1:813"
	}

	Instance = conf

	return nil
}
