package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/carts/domain"
	"github.com/Numsina/tkshop/app/carts/repository"
	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/tools"
)

type CartHandler struct {
	repo repository.Cart
}

func NewCartHandler(repo repository.Cart) *CartHandler {
	return &CartHandler{
		repo: repo,
	}
}

func (c *CartHandler) Create(ctx *gin.Context) {
	var req domain.Carts
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.repo.CreateCarts(ctx.Request.Context(), req, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Update(ctx *gin.Context) {
	var req domain.Carts
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.repo.UpdateCarts(ctx.Request.Context(), req, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Delete(ctx *gin.Context) {
	type Req struct {
		Id int32 `json:"id"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.repo.DeleteCarts(ctx.Request.Context(), req.Id, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Clear(ctx *gin.Context) {
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.repo.ClearCarts(ctx.Request.Context(), claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Get(ctx *gin.Context) {
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	data, err := c.repo.GetCartsInfo(ctx.Request.Context(), claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
	return
}
