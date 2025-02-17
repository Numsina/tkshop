package handler

import (
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"

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
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{
		emailRegexp:    regexp.MustCompile(constant.UserEmail, regexp.None),
		passwordRegexp: regexp.MustCompile(constant.UserPassword, regexp.None),
		birthDayRegexp: regexp.MustCompile(constant.BirthDay, regexp.None),
		NickNameRegexp: regexp.MustCompile(constant.NickName, regexp.None),
		PhoneRegexp:    regexp.MustCompile(constant.PhoneNumber, regexp.None),
		svc:            svc,
	}
}

func (u *UserHandler) singUp(ctx *gin.Context) {
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

	uid, err := u.svc.SignUp(ctx.Request.Context(), user)
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

func (u *UserHandler) GetUserByEmail(ctx *gin.Context) {
	var email = ctx.Query("email")
	if ok, _ := u.emailRegexp.MatchString(email); !ok {
		ctx.JSON(http.StatusOK, "邮箱格式有误!!!")
		return
	}

	user, err := u.svc.GetUserInfoByEmail(ctx.Request.Context(), email)
	if err != nil {
		// 记录日志
		ctx.JSON(http.StatusOK, "内部问题")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "查询成功",
		"data": user,
	})
	return
}
