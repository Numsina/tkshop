package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/tools"
)

func (p *ProductHandler) GetCategoryById(ctx *gin.Context) {
	id := ctx.Query("id")
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

	categorys, err := p.svc.GetCategoryById(ctx.Request.Context(), int32(pid))
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: categorys,
	})
	return

}
func (p *ProductHandler) GetCategoryByName(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	categorys, err := p.svc.GetCategoryByName(ctx.Request.Context(), name)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: categorys,
	})
	return
}
func (p *ProductHandler) CreateCategory(ctx *gin.Context) {
	var req domain.Categorys
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	data, err := p.svc.CreateCategory(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "添加成功",
		Data: data,
	})
	return
}
func (p *ProductHandler) UpdateCategory(ctx *gin.Context) {
	var req domain.Categorys
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	err := p.svc.UpdateCategory(ctx.Request.Context(), req)
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

func (p *ProductHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Query("id")
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

	// claims := ctx.Value("claims").(*middlewares.UserClaims)

	err = p.svc.DeleteCategory(ctx.Request.Context(), int32(pid))
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

// func (p *ProductHandler) GetGetCategoryList(ctx *gin.Context) {

// }
