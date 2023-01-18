package controllers

import (
	"musicAPI/src/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DummyUsers = []entities.User{
	{
		ID:       "192037801297",
		Username: "jepri123",
		Password: "matamu",
	},
	{
		ID:       "102938012938",
		Username: "samsul234",
		Password: "lengser1",
	},
}

func (con *Controller) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := con.Service.GetUser(user.Email, user.Password)
		if err != nil {
			panic(err)
		}
		// for _, v := range DummyUsers {
		// 	if user.Username == v.Username && user.Password == v.Password {
		// 		token, refreshToken, _ := helper.GenerateToken(v.ID)
		// 		c.Request.Header.Set("token", token)
		// 		c.JSON(http.StatusOK, gin.H{
		// 			"token":         token,
		// 			"refresh_token": refreshToken,
		// 			"header_token":  c.Request.Header.Get("token"),
		// 		})
		// 	}
		// }
	}
}
