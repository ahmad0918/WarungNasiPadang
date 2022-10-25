package controller

import (
	"net/http"
	"warung_nasi_padang/delivery/middleware"
	"warung_nasi_padang/model"
	"warung_nasi_padang/usecase"
	"warung_nasi_padang/utils"

	"github.com/gin-gonic/gin"
)

type TransaksiController struct {
	route *gin.Engine
	usecase usecase.TransaksiUsecase
	authUseCase usecase.AuthUseCase
}

func (tc *TransaksiController) createTransaksi(ctx *gin.Context) {
	var transaksi model.Transaksi
	transaksi.Id = utils.UuidGenerate()
	if err := ctx.ShouldBindJSON(&transaksi); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	if err := tc.usecase.CreateTransaksi(transaksi); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transaksi)
}

func (tc *TransaksiController) updateTransaksi(ctx *gin.Context) {
	var transaksi *model.Transaksi
	if err := ctx.ShouldBindJSON(&transaksi); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	if err := tc.usecase.UpdateTransaksi(*transaksi); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transaksi)
}

func (tc *TransaksiController) deleteTransaksi(ctx *gin.Context) {
	idTransaksi := ctx.Param("id")
	if err := tc.usecase.DeleteTransaksi(idTransaksi); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"massage" : "Transaksi berhasil dihapus..!!"})
}

func (tc *TransaksiController) listMenu(ctx *gin.Context) {
	transaksi, err := tc.usecase.ListTransaksi()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err" : err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, transaksi)
}

func NewTransaksiController(r *gin.Engine, usecase usecase.TransaksiUsecase, tokenMdw middleware.AuthTokenMiddleware) *TransaksiController {
	controller := TransaksiController{
		route: r,
		usecase: usecase,
	}
	transaksi := controller.route.Group("/transaksi", tokenMdw.RequireToken())
	transaksi.POST("", controller.createTransaksi)
	transaksi.PUT("", controller.updateTransaksi)
	transaksi.DELETE("/:id", controller.deleteTransaksi)
	transaksi.GET("", controller.listMenu)
	return &controller
}