package database

import (
	"io/ioutil"

	"hako-event-logs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDatabase struct {
	DB *gorm.DB
}

// return new database instance
func NewDatabase(config *config.Config) (*GormDatabase, error) {
	passwordBytes, err := ioutil.ReadFile(config.Mysql.Password)
	if err != nil {
		return nil, err
	}
	userBytes, err := ioutil.ReadFile(config.Mysql.User)
	if err != nil {
		return nil, err
	}
	dsn := string(userBytes) + ":" + string(passwordBytes) + "@/hako?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	return &GormDatabase{
		DB: db,
	}, nil
}

func (d *GormDatabase) Close() {
	db_sql, err := d.DB.DB()
	if err != nil {
		panic(err)
	}
	db_sql.Close()
}
