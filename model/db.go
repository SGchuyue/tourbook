package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"tourbook/utils"
	"time"
	_"github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var err error

func InitDb() {
	db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	if err != nil {
		fmt.Printf("连接数据库发生错误，检查参数：",err)
	}

	// 禁用默认表名的复数模式
	db.SingularTable(true)

	// 将数据模型地址迁移进来
	db.AutoMigrate(&User{},&Article{},&Category{})

	// setMaxIdleConns 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(10)

	// setOpenConns 设置数据库的最大连接数量
	db.DB().SetMaxOpenConns(100)

	// setconnMaxLifetime 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}
