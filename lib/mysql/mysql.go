package mysql

import (
	"fmt"
	"github.com/miko2009/gogogo_follow_me/lib/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hash/crc32"
	"log"
	"os"
	"strconv"
	"time"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

//Mysql 连接池

var Master *gorm.DB
var Slave *gorm.DB


func getMysqlPool(port int, host, user, password, database string) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	var logLevel logger.LogLevel
	if os.Getenv("MY_ENV") == "release" {
		logLevel = 1
	} else {
		logLevel = 4
	}

	mysqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logLevel,    // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		return nil, err
	}
//	metric.AddMySQL(db,dsn)

	pool, err := db.DB()
	// mysql Pool 根据系统状况自行配置
	pool.SetConnMaxLifetime(time.Second * 90)

	pool.SetMaxIdleConns(50)

	pool.SetMaxOpenConns(50)

	err = pool.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

/**
初始化mysql
*/
func InitMysqlPool() error {
	var err error
	conf := config.Config

	Master, err = getMysqlPool(
		conf.Master.Port,
		conf.Master.Host,
		conf.Master.UserName,
		conf.Master.Password,
		conf.Master.DataBase)
	if err != nil {
		return err
	}

	Slave, err = getMysqlPool(
		conf.Slave.Port,
		conf.Slave.Host,
		conf.Slave.UserName,
		conf.Slave.Password,
		conf.Slave.DataBase)
	if err != nil {
		return err
	}


	_ = Master.Callback().Update().Register("updated_at_time_map", updateTimeStampForUpdateCallback)

	return nil
}

/**
分表
*/
func GetSharedTable(sharededId int64, tableNumber int) int {
	idString := strconv.FormatInt(sharededId, 10)
	idUint32 := crc32.ChecksumIEEE([]byte(idString))
	tableIdex := int((int(idUint32) / 100) % tableNumber)
	return tableIdex
}


func updateTimeStampForUpdateCallback(db *gorm.DB) {
	db.Statement.SetColumn("created_at", time.Now().Unix())
}

