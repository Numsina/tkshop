package ioc

import (
	"github.com/Numsina/tkshop/app/user/handler"
	"github.com/Numsina/tkshop/app/user/initialize"
	"github.com/Numsina/tkshop/app/user/repository"
	"github.com/Numsina/tkshop/app/user/repository/dao"
	"github.com/Numsina/tkshop/app/user/service"
)

func InitUser() *handler.UserHandler {
	initialize.InitConfig()
	logger := initialize.InitLogger()
	db := initialize.InitDB()
	dao.InitAutoMigrateTable(db)
	entity := dao.NewUserDao(db, logger)
	userRepo := repository.NewUserRepository(entity)
	userSvc := service.NewUserSvc(userRepo, logger)
	return handler.NewUserHandler(userSvc)
}
