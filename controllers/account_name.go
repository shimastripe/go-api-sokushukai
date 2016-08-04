package controllers

import (
	"net/http"

	dbpkg "github.com/shimastripe/go-api-sokushukai/db"
	"github.com/shimastripe/go-api-sokushukai/helper"
	"github.com/shimastripe/go-api-sokushukai/models"
	"github.com/shimastripe/go-api-sokushukai/version"

	"github.com/gin-gonic/gin"
)

func GetAccountNames(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))

	pagination := dbpkg.Pagination{}
	db, err := pagination.Paginate(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = dbpkg.SetPreloads(preloads, db)

	var accountnames []models.AccountName
	if err := db.Select("*").Find(&accountnames).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// paging
	var index int
	if len(accountnames) < 1 {
		index = 0
	} else {
		index = int(accountnames[len(accountnames)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range(ver, "<", "1.0.0") {
		// conditional branch by version.
		// this version < 1.0.0 !!
		c.JSON(400, gin.H{"error": "this version (< 1.0.0) is not supported!"})
		return
	}

	fieldMap := []map[string]interface{}{}
	for key, _ := range accountnames {
		fieldMap = append(fieldMap, helper.FieldToMap(accountnames[key], fields))
	}
	c.JSON(200, fieldMap)
}

func GetAccountName(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))

	db := dbpkg.DBInstance(c)
	db = dbpkg.SetPreloads(preloads, db)

	var accountname models.AccountName
	if err := db.Select("*").First(&accountname, id).Error; err != nil {
		content := gin.H{"error": "accountname with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range(ver, "<", "1.0.0") {
		// conditional branch by version.
		// this version < 1.0.0 !!
		c.JSON(400, gin.H{"error": "this version (< 1.0.0) is not supported!"})
		return
	}

	fieldMap := helper.FieldToMap(accountname, fields)
	c.JSON(200, fieldMap)
}

func CreateAccountName(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var accountname models.AccountName
	c.Bind(&accountname)
	if db.Create(&accountname).Error != nil {
		content := gin.H{"error": err.Error()}
		c.JSON(500, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, accountname)
}

func UpdateAccountName(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var accountname models.AccountName
	if db.First(&accountname, id).Error != nil {
		content := gin.H{"error": "accountname with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&accountname)
	db.Save(&accountname)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, accountname)
}

func DeleteAccountName(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var accountname models.AccountName
	if db.First(&accountname, id).Error != nil {
		content := gin.H{"error": "accountname with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&accountname)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
