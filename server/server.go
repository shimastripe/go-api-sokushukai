package server

import (
	"github.com/shimastripe/go-api-sokushukai/middleware"
	"github.com/shimastripe/go-api-sokushukai/router"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Setup(db *gorm.DB) *gin.Engine {
  r := gin.Default()
  r.Use(middleware.SetDBtoContext(db))
  router.Initialize(r)
  return r
}
