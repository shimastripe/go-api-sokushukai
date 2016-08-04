package controllers

import (
	"net/http"

	dbpkg "github.com/shimastripe/go-api-sokushukai/db"
	"github.com/shimastripe/go-api-sokushukai/helper"
	"github.com/shimastripe/go-api-sokushukai/models"
	"github.com/shimastripe/go-api-sokushukai/version"

	"github.com/gin-gonic/gin"
)

func GetEmails(c *gin.Context) {
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

	var emails []models.Email
	if err := db.Select("*").Find(&emails).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// paging
	var index int
	if len(emails) < 1 {
		index = 0
	} else {
		index = int(emails[len(emails)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range(ver, "<", "1.0.0") {
		// conditional branch by version.
		// this version < 1.0.0 !!
		c.JSON(400, gin.H{"error": "this version (< 1.0.0) is not supported!"})
		return
	}

	fieldMap := []map[string]interface{}{}
	for key, _ := range emails {
		fieldMap = append(fieldMap, helper.FieldToMap(emails[key], fields))
	}
	c.JSON(200, fieldMap)
}

func GetEmail(c *gin.Context) {
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

	var email models.Email
	if err := db.Select("*").First(&email, id).Error; err != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range(ver, "<", "1.0.0") {
		// conditional branch by version.
		// this version < 1.0.0 !!
		c.JSON(400, gin.H{"error": "this version (< 1.0.0) is not supported!"})
		return
	}

	fieldMap := helper.FieldToMap(email, fields)
	c.JSON(200, fieldMap)
}

func CreateEmail(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var email models.Email
	c.Bind(&email)
	if db.Create(&email).Error != nil {
		content := gin.H{"error": err.Error()}
		c.JSON(500, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, email)
}

func UpdateEmail(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var email models.Email
	if db.First(&email, id).Error != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&email)
	db.Save(&email)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, email)
}

func DeleteEmail(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var email models.Email
	if db.First(&email, id).Error != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&email)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
