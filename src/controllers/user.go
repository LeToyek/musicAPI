package controllers

import (
	"musicAPI/src/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
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
