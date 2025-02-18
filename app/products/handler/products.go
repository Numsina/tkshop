package handler

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/app/products/service"
)

type ProductHandler struct {
	svc    service.Product
	logger *zap.Logger
}

func NewProductHandler(svc service.Product, logger *zap.Logger) *ProductHandler {
	return &ProductHandler{
		svc:    svc,
		logger: logger,
	}
}

func (p *ProductHandler) Create(ctx *gin.Context) {
	var pd domain.Products
	if err := ctx.Bind(&pd); err != nil {
		p.logger.Sugar().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	products, err := p.svc.CraeteProduct(ctx.Request.Context(), pd)
	if err != nil {
		ctx.JSON(http.StatusOK, "创建失败")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "商品创建成功",
		"data": products,
	})
	return
}

func (p *ProductHandler) Update(ctx *gin.Context) {
	var pd domain.Products
	if err := ctx.Bind(&pd); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	products, err := p.svc.UpdateProduct(ctx.Request.Context(), pd)
	if err != nil {
		ctx.JSON(http.StatusOK, "更新商品失败")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "更新商品成功",
		"data": products,
	})
	return
}

func (p *ProductHandler) GetProductsDetail(ctx *gin.Context) {
	var id = ctx.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}
	if pid == 0 {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	products, err := p.svc.GetProductInfoById(ctx.Request.Context(), int32(pid))
	if err != nil {
		ctx.JSON(http.StatusOK, "获取商品详情失败")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "获取商品详情成功",
		"data": products,
	})
	return
}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	type delete_req struct {
		Id int32 `json:"id"`
	}
	var req delete_req
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	err := p.svc.DeleteProduct(ctx.Request.Context(), int32(req.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, "商品删除失败")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "商品删除成功",
	})
	return
}
