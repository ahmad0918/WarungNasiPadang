package controller

import (
	"net/http"
	"warung_nasi_padang/delivery/middleware"
	"warung_nasi_padang/model"
	"warung_nasi_padang/usecase"
	"warung_nasi_padang/utils"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	route *gin.Engine
	usecase usecase.MenuUsecase
	authUseCase usecase.AuthUseCase
}

func (mc *MenuController) userAuth(ctx *gin.Context) {
	var user model.UserCredential
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message":"can't bind struct",
		})
		return
	}
	token, err := mc.authUseCase.UserAuth(user)
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"token":token})
}

func (mc *MenuController) createMenu(ctx *gin.Context) {
	var menu model.Menu
	menu.Id = utils.UuidGenerate()
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	if err := mc.usecase.CreateMenu(menu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func (mc *MenuController) updateMenu(ctx *gin.Context) {
	var menu *model.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	if err := mc.usecase.UpdateMenu(*menu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, menu)	
}

func (mc *MenuController) deleteMenu(ctx *gin.Context) {
	idMenu := ctx.Param("id")
	if err := mc.usecase.DeleteMenu(idMenu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"massage" : "Menu berhasil dihapus..!!"})
}

func (mc *MenuController) listMenu(ctx *gin.Context) {
	menu ,err := mc.usecase.ListMenu()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, menu)
}

func NewMenuController(r *gin.Engine, usecase usecase.MenuUsecase, auth usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleware) *MenuController {
	controller := MenuController{
		route: r,
		usecase: usecase,
		authUseCase: auth,
	}
	
	controller.route.POST("/auth",controller.userAuth)
	menu := controller.route.Group("/menu",tokenMdw.RequireToken())
	menu.POST("", controller.createMenu)
	menu.PUT("", controller.updateMenu)
	menu.DELETE("/:id", controller.deleteMenu)
	menu.GET("",controller.listMenu)
	return &controller
}