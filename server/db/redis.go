package db

import (
	"errors"
	"github.com/go-redis/redis"
	"net"
	"time"
	"wms/config"
	"wms/log"
)

var Redis *redis.Client

func InitRedis() error {
	log.Infof("================[开始初始化Redis数据库连接]================")
	_, _, err := net.SplitHostPort(config.Instance.Redis.Address)
	if err != nil {
		return err
	}
	opt := &redis.Options{
		Addr:         config.Instance.Redis.Address,
		MaxRetries:   3,
		DB:           config.Instance.Redis.Database,
		DialTimeout:  200 * time.Millisecond,
		ReadTimeout:  350 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		PoolSize:     1000,
	}
	if config.Instance.Redis.Password != "" {
		opt.Password = config.Instance.Redis.Password
	}
	client := redis.NewClient(opt)
	if client == nil {
		log.Error("InitRedis error")
		return errors.New("init client failed")
	}
	Redis = client
	log.Infof("================[结束初始化Redis数据库连接]================")
	return nil
}
