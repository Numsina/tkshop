package handler

import (
	"log"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/Numsina/tkshop/app/middlewares"
	"github.com/Numsina/tkshop/app/user/constant"
	"github.com/Numsina/tkshop/app/user/domain"
	"github.com/Numsina/tkshop/app/user/service"
)

type UserHandler struct {
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
	birthDayRegexp *regexp.Regexp
	NickNameRegexp *regexp.Regexp
	PhoneRegexp    *regexp.Regexp
	svc            service.UserService
	jhl            *middlewares.JWT
	logger         *zap.Logger
}

func NewUserHandler(svc service.UserService, jhl *middlewares.JWT, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		emailRegexp:    regexp.MustCompile(constant.UserEmail, regexp.None),
		passwordRegexp: regexp.MustCompile(constant.UserPassword, regexp.None),
		birthDayRegexp: regexp.MustCompile(constant.BirthDay, regexp.None),
		NickNameRegexp: regexp.MustCompile(constant.NickName, regexp.None),
		PhoneRegexp:    regexp.MustCompile(constant.PhoneNumber, regexp.None),
		svc:            svc,
		jhl:            jhl,
		logger:         logger,
	}
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	var user domain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	if ok, _ := u.emailRegexp.MatchString(user.Email); !ok {
		ctx.JSON(http.StatusOK, "邮箱格式有误!!!")
		return
	}

	if ok, _ := u.passwordRegexp.MatchString(user.Password); !ok {
		ctx.JSON(http.StatusOK, "密码格式有误!!!")
		return
	}

	if user.Password != user.ConfirmPassword {
		ctx.JSON(http.StatusOK, "两次输入密码不一致!!!")
		return
	}

	uid, err := u.svc.SignUp(ctx.Request.Context(), domain.User{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		// 可能是数据库的错误， 可能是加密的错误， 可能唯一主键的错误， 记录日志，
		// 如果唯一主键错误，则返回已创建，否则内部错误
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "用户创建成功",
		"data": uid,
	})
	return
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type login_req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req login_req
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	if ok, _ := u.emailRegexp.MatchString(req.Email); !ok {
		ctx.JSON(http.StatusOK, "邮箱或密码不正确")
		return
	}

	if ok, _ := u.passwordRegexp.MatchString(req.Password); !ok {
		ctx.JSON(http.StatusOK, "邮箱或密码不正确")
		return
	}

	ue, err := u.svc.Login(ctx.Request.Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, "登陆失败")
		return
	}

	// 设置token或者cookie
	// 生成session，并设置token
	uid := uuid.New()
	tokenString, err := u.jhl.SetToken(ctx, ue.Id, uid.String())
	if err != nil {
		u.logger.Error("生成token失败")
		ctx.JSON(http.StatusOK, "登录失败")
		return
	}

	ctx.Header("x-jwt-token", tokenString)

	ctx.JSON(http.StatusOK, "登陆成功")
	return
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	log.Println(claims)
	err := u.jhl.DeleteSsid(ctx, claims)
	if err != nil {
		u.logger.Info("删除ssid失败")
		ctx.JSON(http.StatusOK, "退出登录失败")
		return
	}
	ctx.JSON(http.StatusOK, "退出成功")
	return
}

func (u *UserHandler) Update(ctx *gin.Context) {
	var user domain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	if ok, _ := u.emailRegexp.MatchString(user.Email); !ok {
		ctx.JSON(http.StatusOK, "邮箱格式有误!!!")
		return
	}

	if user.Password != user.ConfirmPassword {
		ctx.JSON(http.StatusOK, "两次输入密码不一致!!!")
		return
	}

	user, err := u.svc.ModifyUserInfoById(ctx.Request.Context(), user)
	if err != nil {
		// 记录日志
		ctx.JSON(http.StatusOK, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "修改成功",
		"data": user,
	})
	return
}

func (u *UserHandler) Delete(ctx *gin.Context) {
	type delete_req struct {
		Uid int32 `json:"uid"`
	}
	var req delete_req
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}

	// TODO
	val, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, "身份未授权")
		return
	}
	uid := val.(int32)

	if uid != req.Uid {
		ctx.JSON(http.StatusUnauthorized, "身份未授权")
		return
	}
	err := u.svc.Delele(ctx.Request.Context(), req.Uid)

	if err != nil {
		ctx.JSON(http.StatusOK, "删除失败")
		return
	}

	ctx.JSON(http.StatusOK, "删除成功")
	return
}

func (u *UserHandler) 	GetUserByEmail(ctx *gin.Context) {
	var email = ctx.Query("email")
	if ok, _ := u.emailRegexp.MatchString(email); !ok {
		ctx.JSON(http.StatusOK, "邮箱格式有误!!!")
		return
	}

	user, err := u.svc.GetUserInfoByEmail(ctx.Request.Context(), email)
	if err != nil {
		ctx.JSON(http.StatusOK, "内部问题")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "查询成功",
		"data": user,
	})
	return
}
