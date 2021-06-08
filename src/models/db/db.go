package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"web-project-model/src/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//DB 数据库db
var DB *gorm.DB

func DBInit() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       utils.GetDBConfig(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		err = fmt.Errorf("[DB] gorm open mysql failed: %s", err.Error())
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		err = fmt.Errorf("[DB] gorm connect the pool failed: %s", err.Error())
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)
	// Set table options
	DB.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Set("gorm:table_options", "ENGINE=InnoDB")
	err = DB.Set("gorm:table_options", "CHARSET=utf8").Set("gorm:table_options", "COLLATE=utf8_general_ci").AutoMigrate(
		&InterestsOrder{},
	)
	if err != nil {
		err = fmt.Errorf("[DB] create table failed: %s", err.Error())
		panic(err)
	}
}
