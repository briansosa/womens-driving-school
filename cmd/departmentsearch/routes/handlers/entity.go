package handlers

import (
	"braiton/braiton-home/cmd/internal/application"
	"braiton/braiton-home/cmd/internal/domain"
	"braiton/braiton-home/cmd/internal/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEntities(c *gin.Context) {
	relationship, _ := strconv.ParseBool(c.Query("withRelationship"))
	var entities []domain.Entity
	var err error
	if relationship {
		entities, err = application.GetEntitiesWithRelationship()
	} else {
		entities, err = application.GetEntities()
	}
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	var entitiesDto []domain.EntityDto
	err = utils.MapperStructs(entities, &entitiesDto)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, entitiesDto)
}

func GetEntity(c *gin.Context) {
	relationship, _ := strconv.ParseBool(c.Query("withRelationship"))
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	var entity domain.Entity
	var err error
	if relationship {
		entity, err = application.GetEntityWithRelationship(uint(id))
	} else {
		entity, err = application.GetEntity(uint(id))
	}

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	var entityDto domain.EntityDto
	err = utils.MapperStructs(entity, &entityDto)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, entityDto)
}

func CreateEntity(c *gin.Context) {
	var entityDto domain.EntityDto
	c.BindJSON(&entityDto)

	err := application.CreateEntity(entityDto)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, entityDto)
}

func UpdateEntity(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	var entityDto domain.EntityDto
	c.BindJSON(&entityDto)

	err := application.UpdateEntity(entityDto, uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, entityDto)
}

func DeleteEntity(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	err := application.DeleteEntity(uint(id))
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, id)
}
