package utils

import (
	"fmt"
	"ginchat/common"
	"ginchat/model"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"os"
	"time"
)

// viper设置配置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("初始化成功！")
}

// zap日志库的初始化

func GetLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 只打印在控制台
// func InitLogger() {
//    writeSyncer := getLogWriter()
//    encoder := getEncoder()
//    core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//
//    logger := zap.New(core, zap.AddCaller())
//    model.common.SugarLogger = logger.Sugar()
// }

func InitLogger() {
	// 打印到日志和终端
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel)

	writeSyncer := zapcore.NewMultiWriteSyncer(consoleDebugging, GetLogWriter())
	encoder := GetEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(zapcore.NewTee(consoleCore, core), zap.AddCaller())
	common.SugarLogger = logger.Sugar()
	defer common.SugarLogger.Sync()
	SimpleHttpGet("www.sogo.com")
	SimpleHttpGet("http://www.sogo.com")
	//SimpleHttpGet("http://localhost:6060/")

}

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func SimpleHttpGet(url string) {
	common.SugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		common.SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		common.SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

// mysql初始化
func InitMysql() {

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{
			Logger: logger.New(
				zap.NewStdLog(common.SugarLogger.Desugar()), //设置zap日志记录器
				logger.Config{
					SlowThreshold: time.Second, //慢SQL阈值
					LogLevel:      logger.Info, //级别
					Colorful:      true,        //彩色
				},
			),
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	// 迁移 schema
	db.AutoMigrate(&model.User{})
	common.DB = db
}

// redis初始化
func InitRedis() {
	common.RDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
}
