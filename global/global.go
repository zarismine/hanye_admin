package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Config     *viper.Viper
	ShopStatus int
	// Redis *redis.Client
)
