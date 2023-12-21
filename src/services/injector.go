package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	GormDBInstance = "gorm_db_instance"
	ConfigInstance = "config_instance"
)

func InjectGormDB(ctx *gin.Context, db *gorm.DB) {
	ctx.Set(GormDBInstance, db)
}

func GetGormDBInstance(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(GormDBInstance).(*gorm.DB)
}

func InjectConfig(ctx *gin.Context, cnf *Config) {
	ctx.Set(ConfigInstance, cnf)
}

func GetConfigInstance(ctx *gin.Context) *Config {
	return ctx.MustGet(ConfigInstance).(*Config)
}
