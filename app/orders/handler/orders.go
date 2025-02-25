package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/orders/domain"
	"github.com/Numsina/tkshop/app/orders/repository"
	"github.com/Numsina/tkshop/tools"
)

type OrderHandler struct {
	repo repository.Order
}

func NewOrderHandler(repo repository.Order) *OrderHandler {
	return &OrderHandler{
		repo: repo,
	}
}

func (o *OrderHandler) Create(ctx *gin.Context) {
	type Req struct {
		PayType string `json:"pay_type"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := o.repo.SaveOrder(ctx.Request.Context(), domain.Orders{
		PayType: req.PayType,
		Address: req.Address,
		Phone:   req.Phone,
	}, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 应该返回支付链接
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
		Data: "",
	})
	return
}

//func (o *OrderHandler) Update(ctx *gin.Context) {
//
//}

func (o *OrderHandler) GetByUid(ctx *gin.Context) {
	type Req struct {
		Page int32 `json:"page"`
		Size int32 `json:"size"`
	}
	var req Req
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	data, err := o.repo.GetOrderByUid(ctx.Request.Context(), claims.UserId, req.Page, req.Size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, tools.Result{
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

func (o *OrderHandler) Update(ctx *gin.Context) {

}

func (o *OrderHandler) Delete(ctx *gin.Context) {
	type Req struct {
		Sn string `json:"sn"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := o.repo.DeleteOrder(ctx, req.Sn, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, tools.Result{
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

func (o *OrderHandler) Search(ctx *gin.Context) {

	type Req struct {
		OrderSn string `json:"order_sn"`
		PayType string `json:"pay_type"`
		Status  int32  `json:"status"`
		PayTime int64  `json:"pay_time"`
		Page    int32  `json:"page"`
		Size    int32  `json:"size"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	order, err := o.repo.GetOrder(ctx.Request.Context(), domain.Orders{

		OrderSn: req.OrderSn,
		PayType: req.PayType,
		Status:  req.Status,
		PayTime: req.PayTime,
	}, int(req.Page), int(req.Size))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, tools.Result{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
		Data: order,
	})
	return
}
