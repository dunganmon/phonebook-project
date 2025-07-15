package middlewares

import (
	"TODO/databases"
	"TODO/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CheckOwnContact() gin.HandlerFunc {
	return func(c *gin.Context) {
		contactIdStr := c.Param("id")
		contactIdInt, err := strconv.Atoi(contactIdStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "fail to convert string to int",
			})
			return
		}

		contactId := uint(contactIdInt)

		user := c.MustGet("user").(*models.User)

		fmt.Println("UserId: ", user.ID, " - contactId: ", contactId)
		contact := models.Contact{}
		err = databases.DB.Where("id=? AND user_id=?", contactId, user.ID).First(&contact).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "fail to delete contact",
			})
			return
		}

		if contact.ID != contactId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "fail to delete contact",
			})
			return
		}

		c.Next()

	}
}
