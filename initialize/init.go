package initialize

import (
	"admin_app/global"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() {
	global.Config = viper.New()
	global.Config.SetConfigName("admin.env")
	global.Config.SetConfigType("yaml")
	global.Config.AddConfigPath("config")
	err := global.Config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config mysql:", global.Config.Get("mysql"))
}

func Init_DB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(global.Config.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" MySQL inited.......")
}
func Init() {
	InitConfig()
	Init_DB()
}
