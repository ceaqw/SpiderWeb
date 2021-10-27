package conf

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

// 开发环境
const BUILD = "dev"

// 生产环境
// const BUILD = "prod"

func init() {
	workPath, _ := os.Getwd()
	viper.SetConfigName("service_" + BUILD)
	viper.SetConfigType("yml")
	viper.AddConfigPath(workPath + "/conf")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

type MainDbCfg struct {
	UserName string
	Password string
	Database string
	Host     string
	Port     string
	Driver   string
}

type AppCfg struct {
	Port string //App 启动端口
	Host string //App
}

type JwtCfg struct {
	TimeOut time.Duration
	Issuer  string
}

type LoggerCfg struct {
	LogPath string
	LogName string
}

func GetMainDbCfg() (mainDbCfg MainDbCfg) {
	mainDbCfg.UserName = viper.GetString("models.main_db.username")
	mainDbCfg.Password = viper.GetString("models.main_db.password")
	mainDbCfg.Database = viper.GetString("models.main_db.database")
	mainDbCfg.Host = viper.GetString("models.main_db.host")
	mainDbCfg.Port = viper.GetString("models.main_db.port")
	mainDbCfg.Driver = viper.GetString("models.main_db.driver")
	return mainDbCfg
}

func GetAppCfg() (appCfg AppCfg) {
	appCfg.Port = viper.GetString("service.port")
	appCfg.Host = viper.GetString("service.host")
	return appCfg
}

func GetJwtCfg() (jwtCfg JwtCfg) {
	jwtCfg.TimeOut = viper.GetDuration("jwt.time_out")
	jwtCfg.Issuer = viper.GetString("jwt.issuer")
	return jwtCfg
}

func GetLoggerCfg() (loggerCfg LoggerCfg) {
	loggerCfg.LogPath = viper.GetString("logger.log_path")
	loggerCfg.LogName = viper.GetString("logger.log_name")
	return loggerCfg
}
