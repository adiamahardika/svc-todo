package router

import (
	"os"
	"svc-todo/controller"
	"svc-todo/customValidator"
	"svc-todo/repository"
	"svc-todo/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {

	router := echo.New()
	router.Validator = &customValidator.CustomValidator{Validator: validator.New()}

	repository := repository.Repository(db)

	listService := service.ListService(repository)
	listController := controller.ListController(listService)

	subListService := service.SubListService(repository)
	subListController := controller.SubListController(subListService)

	dir := os.Getenv("FILE_DIR")
	router.Static("/assets", dir)

	todo := router.Group("/todo")
	{
		v1 := todo.Group("/v1")
		{
			list := v1.Group("/list")
			{
				list.POST("/create", listController.CreateList)
				list.GET("/get", listController.GetList)
				list.GET("/get-with-sub", listController.GetListWithSub)
			}

			sub_list := v1.Group("/sub-list")
			{
				sub_list.POST("/create", subListController.CreateSubList)
				sub_list.GET("/get", subListController.GetSubList)
			}
		}
	}

	router.Start(os.Getenv("PORT"))
}
