//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"orenotorero/handler"
	"orenotorero/infrastructure/mysqlDB"
	"orenotorero/service"
)

func InitUserAPI(db *gorm.DB) handler.UserHandler {
	wire.Build(mysqlDB.NewUserRepoImpl, service.NewUserService, handler.NewUserHandler)

	return handler.UserHandler{}
}

func InitUtilityAPI(db *gorm.DB) handler.UtilityHandler {
	wire.Build(mysqlDB.NewUserRepoImpl, mysqlDB.NewKanbanRepoImpl, service.NewUtilityService, handler.NewUtilityHandler)

	return handler.UtilityHandler{}
}

func InitBoardAPI(db *gorm.DB) handler.BoardHandler {
	wire.Build(mysqlDB.NewBoardRepoImpl, service.NewBoardService, handler.NewBoardHandler)

	return handler.BoardHandler{}
}

func InitCardAPI(db *gorm.DB) handler.CardHandler {
	wire.Build(mysqlDB.NewCardRepoImpl, mysqlDB.NewFileRepoImpl, service.NewCardService, handler.NewCardHandler)

	return handler.CardHandler{}
}

func InitKanbanAPI(db *gorm.DB) handler.KanbanHandler {
	wire.Build(mysqlDB.NewKanbanRepoImpl, service.NewKanbanService, handler.NewKanbanHandler)

	return handler.KanbanHandler{}
}
