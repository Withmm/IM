package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Rds *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited")
	fmt.Println("config mysql inited")
}

func InitMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	fmt.Println(viper.GetString("mysql.dns"))
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to databases")
	}
	//user := &models.UserBasic{}
	//db.Find(user)
	//fmt.Println(user)
}
func InitRedis() {
	Rds = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish...", msg)
	err = Rds.Publish(ctx, channel, msg).Err()
	return err
}

// Subscibe 订阅redis消息
func Subscibe(ctx context.Context, channel string) (string, error) {
	sub := Rds.Subscribe(ctx, channel)
	fmt.Println("1:Subscribe...", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("2:Subscribe...", msg.Payload)
	return msg.Payload, err
}
