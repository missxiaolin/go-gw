package database

import (
	"github.com/jinzhu/gorm"
	"go-gw/config"
)

var DB *gorm.DB

func init()  {
	cfg := config.Cfg
	optionStr := cfg.Mysql.SqlUsername + ":" + cfg.Mysql.SqlPassword + "@tcp(" + cfg.Mysql.SqlHost + ":" + cfg.Mysql.SqlPort + ")/" + cfg.Mysql.SqlDb + "?charset=utf8mb4&parseTime=True&loc=Local"
}