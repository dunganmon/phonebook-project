package controllers

import (
	"TODO/dtos"
	"TODO/models"
	"TODO/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetContacts(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	contacts, err := services.GetContactsByUserId(int(user.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get contacts successfully",
		"data": gin.H{
			"contacts": contacts,
		},
	})
}

func AddContact(c *gin.Context) {
	var contactReq dtos.ContactRequest
	if err := c.ShouldBindJSON(&contactReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*models.User)
	contactReq.UserID = int(user.ID)
	fmt.Println(contactReq.UserID)

	contact, err := services.AddContact(contactReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Contact added successfully",
		"data": gin.H{
			"contact": contact,
		},
	})
}

func DeleteContact(c *gin.Context) {
	contactIdStr := c.Param("id")
	contactId, err := strconv.Atoi(contactIdStr)

	err = services.DeleteContact(contactId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Contact deleted successfully",
		"data": gin.H{
			"contactId": contactId,
		},
	})
}

func EditContact(c *gin.Context) {
	contactIdStr := c.Param("id")
	contactId, err := strconv.Atoi(contactIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var contactReq dtos.ContactRequest
	if err := c.ShouldBindJSON(&contactReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	contact, errService := services.EditContact(contactId, contactReq)
	if errService != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errService.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Contact edited successfully",
		"data": gin.H{
			"contactId": contact,
		},
	})
}
