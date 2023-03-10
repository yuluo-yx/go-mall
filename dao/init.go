package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

// 初始化连接数据库配置

var _db *gorm.DB

func Database(connRead, connWrite string) {
	// 日志
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		// 如果是debug模式 则 打印日志信息
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		// mysql配置
		DSN:                       connRead,
		DefaultStringSize:         256,  // string类型字段默认长度
		DisableDatetimePrecision:  true, // 禁止datetime精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引，就要把索引先删除之后在重新创建，mysql5.7不支持
		DontSupportRenameColumn:   true, // 用change重命名时，mysql8之前的版本不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		// gorm 配置
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  // 设置连接池
	sqlDB.SetMaxOpenConns(100) // 设置打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// mysql主从配置 使用gorm的一个插件
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		// 写操作配置
		Sources: []gorm.Dialector{mysql.Open(connWrite)},
		// 读操作配置
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)},
		// 策略配置
		Policy: dbresolver.RandomPolicy{},
	}))

	// 自动迁移
	Migration()

}

// NewDBClient 创建一个数据库服务客户端
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
