package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/tools"
)

func (p *ProductHandler) GetBrandById(ctx *gin.Context) {
	id := ctx.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	if pid == 0 {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	brands, err := p.svc.GetBrandById(ctx.Request.Context(), int32(pid))
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	brands.Uid = 0
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: brands,
	})
	return
}

func (p *ProductHandler) GetBrandByName(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	brands, err := p.svc.GetBrandByName(ctx.Request.Context(), name)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	brands.Uid = 0
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: brands,
	})
	return
}

func (p *ProductHandler) CreateBrand(ctx *gin.Context) {
	var req domain.Brands
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	if req.Uid != 0 && req.Uid != claims.UserId {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
		})
		return
	}
	req.Uid = claims.UserId
	data, err := p.svc.CreateBrand(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	data.Uid = 0
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "添加成功",
		Data: data,
	})
	return
}

func (p *ProductHandler) UpdateBrand(ctx *gin.Context) {
	var req domain.Brands
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	if req.Uid != 0 && req.Uid != claims.UserId {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
		})
		return
	}
	req.Uid = claims.UserId
	err := p.svc.UpdateBrand(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "修改成功",
	})
	return
}

func (p *ProductHandler) DeleteBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	if pid == 0 {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)

	err = p.svc.DeleteBrand(ctx.Request.Context(), int32(pid), claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "删除成功",
	})
	return
}

func (p *ProductHandler) GetBrandList(ctx *gin.Context) {
	type Req struct {
		Num  int32 `json:"num"`
		Size int32 `json:"size"`
	}

	var req Req
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	data, err := p.svc.GetBrandList(ctx, req.Num, req.Size)

	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	for _, brand := range data {
		brand.Uid = 0
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Data: data,
	})
	return
}

func (p *ProductHandler) GetBrandByUid(ctx *gin.Context) {
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	data, err := p.svc.GetBrandByUid(ctx, claims.UserId)

	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	for _, brand := range data {
		brand.Uid = 0
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Data: data,
	})
	return
}
