package controller

import (
	"net/http"
	"svc-todo/general"
	"svc-todo/model"
	"svc-todo/service"

	"github.com/labstack/echo"
)

type listController struct {
	listService service.ListServiceInterface
}

func ListController(listService service.ListServiceInterface) *listController {
	return &listController{listService}
}

func (controller *listController) CreateList(context echo.Context) error {

	request := new(model.CreateListRequest)

	error := context.Bind(request)
	if error == nil {
		error = context.Validate(request)
	}
	description := []string{}
	var status *model.StatusResponse

	if error != nil {

		description = append(description, error.Error())

		status = &model.StatusResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}
		return context.JSON(http.StatusBadRequest, model.StandardResponse{
			Status: *status,
		})
	} else {
		list, error := controller.listService.CreateList(request, context)

		if error == nil {

			description = append(description, "Success")

			status = &model.StatusResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			return context.JSON(http.StatusOK, model.StandardResponse{
				Status: *status,
				Result: list,
			})

		} else {

			description = append(description, error.Error())

			status = &model.StatusResponse{
				HttpStatusCode: http.StatusBadRequest,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}
			return context.JSON(http.StatusBadRequest, model.StandardResponse{
				Status: *status,
			})

		}
	}
}

func (controller *listController) GetList(context echo.Context) error {

	request := new(model.GetListRequest)
	error := context.Bind(request)

	description := []string{}
	var status *model.StatusResponse

	if error != nil {

		description = append(description, error.Error())

		status = &model.StatusResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}
		return context.JSON(http.StatusBadRequest, model.StandardResponse{
			Status: *status,
		})
	} else {
		list, total_pages, error := controller.listService.GetList(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StatusResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}
			page := struct {
				Page       int `json:"page"`
				TotalPages int `json:"totalPages"`
			}{
				Page:       request.PageNo,
				TotalPages: total_pages,
			}
			return context.JSON(http.StatusOK, model.StandardResponse{
				Status: *status,
				Result: list,
				Page:   page,
			})

		} else {

			description = append(description, error.Error())

			status = &model.StatusResponse{
				HttpStatusCode: http.StatusBadRequest,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}
			return context.JSON(http.StatusBadRequest, model.StandardResponse{
				Status: *status,
			})

		}
	}
}
