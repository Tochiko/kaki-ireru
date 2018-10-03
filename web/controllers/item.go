package controllers

import (
	"github.com/gin-gonic/gin"
	"kaki-ireru/web/models"
	"net/http"
	"strconv"
)

func (env *Env) AllItems (c *gin.Context) {
	user := &models.User{Id: c.GetString("userId")}
	items, err := env.Db.AllItems(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (env *Env) ItemById (c *gin.Context) {
	user := &models.User{Id: c.GetString("userId")}
	itemId, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "the id of an item must be an int"})
		return
	}
	item, err := env.Db.ItemById(user, itemId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (env *Env) CreateItem (c *gin.Context) {
	user := &models.User{Id: c.GetString("userId")}
	var item models.Item
	c.BindJSON(&item)

	_, err := env.Db.CreateItem(user, &item)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}