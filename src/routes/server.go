package routes

import (
	"musicAPI/src/controllers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router     *gin.Engine
	Controller *controllers.Controller
}
