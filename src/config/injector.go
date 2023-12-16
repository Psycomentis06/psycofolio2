package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	GormDBInstance = "gorm_db_instance"
)

func InjectGormDB(ctx *gin.Context, db *gorm.DB) {
	ctx.Set(GormDBInstance, db)
}

func GetGormDBInstance(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(GormDBInstance).(*gorm.DB)
}
