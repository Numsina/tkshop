package handler

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Numsina/tkshop/app/products/domain"
	"github.com/Numsina/tkshop/app/products/service"
	"github.com/Numsina/tkshop/tools"
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
	var id = ctx.Query("id")
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

	// 应该增加该商品的浏览量, 起一个协程去做, 成不成功无所谓（应该建一个表，用来显示用户浏览过那些商品和收藏过那些）
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err := p.svc.IncreateClick(ctx.Request.Context(), products.Id)
		if err != nil {
			p.logger.Info("增加浏览量失败")
		}
		wg.Done()
	}()

	wg.Wait()
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

func (p *ProductHandler) GetProductsList(ctx *gin.Context) {
	type list_req struct {
		Name         string  `json:"name"`
		CategoryId   int32   `json:"category_id"`
		CategoryName string  `json:"category_name"`
		BrandId      int32   `json:"brand_id"`
		BrandName    string  `json:"brand_name"`
		IsNew        bool    `json:"is_new"`
		IsHot        bool    `json:"is_hot"`
		OnSale       bool    `json:"on_sale"`
		Sale         int32   `json:"sale"`
		MarkPrice    float32 `json:"mark_price"`
		ShopPrice    float32 `json:"shop_price"`
		PNum         int32   `json:"pNum"`
		PSize        int32   `json:"pSize"`
	}
	var data list_req
	if err := ctx.Bind(&data); err != nil {
		p.logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, tools.Result{
			// 错误码待设计
			// todo
			Code: 0,
			Msg:  "参数错误",
		})
		return
	}

	products, err := p.svc.GetProductList(ctx.Request.Context(), domain.Products{
		Name:         data.Name,
		BrandId:      data.BrandId,
		BrandName:    data.BrandName,
		CategoryId:   data.CategoryId,
		CategoryName: data.CategoryName,
		IsNew:        data.IsNew,
		IsHot:        data.IsHot,
		Sale:         data.Sale,
		OnSale:       data.OnSale,
		MarkPrice:    data.MarkPrice,
		ShopPrice:    data.ShopPrice,
	}, data.PNum, data.PSize)

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
		Data: products,
	})
	return
}

func (p *ProductHandler) AddFavorite(ctx *gin.Context) {
	type Req struct {
		Id int32 `json:"id"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		p.logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	// todo 待修改
	// claims := ctx.Value("claims").(*middlewares.UserClaims)

	err := p.svc.IncreateFavorite(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, tools.Result{
			Code: -1,
			Msg:  "添加收藏失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "添加收藏成功",
	})
	return
}
