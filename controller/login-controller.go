package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jtjin/golang-pratice/entity"
	"github.com/jtjin/golang-pratice/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func NewLoginController(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var user entity.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(user.Account, user.Password, user.Company)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(user.Account, user.Company)
	}
	return ""
}
