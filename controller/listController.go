package controller

import (
	"net/http"
	"svc-todo/entity"
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
	var http_status int
	var status *model.StatusResponse

	if error != nil {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest
		status = &model.StatusResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}

	} else {
		_, error = controller.listService.CreateList(request, context)

		if error == nil {

			description = append(description, "Success")
			http_status = http.StatusOK
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusBadRequest,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}
		}
	}

	return context.JSON(http_status, model.StandardResponse{
		Status: *status,
	})
}

func (controller *listController) GetList(context echo.Context) error {

	request := new(model.GetListRequest)
	error := context.Bind(request)
	description := []string{}
	var http_status int
	var status *model.StatusResponse
	var list []entity.List
	var total_pages int

	if error != nil {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest
		status = &model.StatusResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}

	} else {
		list, total_pages, error = controller.listService.GetList(request)

		if error == nil {

			description = append(description, "Success")
			http_status = http.StatusOK
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusBadRequest,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}

		}
	}

	page := model.PaginationResponse{
		Page:       request.PageNo,
		TotalPages: total_pages,
	}
	return context.JSON(http_status, model.StandardResponse{
		Status: *status,
		Result: list,
		Page:   page,
	})
}

func (controller *listController) GetListWithSub(context echo.Context) error {

	request := new(model.GetListRequest)
	error := context.Bind(request)

	description := []string{}
	var http_status int
	var status *model.StatusResponse
	var list []model.GetListResponse
	var total_pages int

	if error != nil {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest
		status = &model.StatusResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}

	} else {
		list, total_pages, error = controller.listService.GetListWithSub(request)

		if error == nil {

			description = append(description, "Success")
			http_status = http.StatusOK
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusBadRequest,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}

		}
	}

	page := model.PaginationResponse{
		Page:       request.PageNo,
		TotalPages: total_pages,
	}

	return context.JSON(http_status, model.StandardResponse{
		Status: *status,
		Result: list,
		Page:   page,
	})
}

func (controller *listController) UpdateList(context echo.Context) error {

	request := new(model.UpdateListRequest)

	error := context.Bind(request)
	if error == nil {
		error = context.Validate(request)
	}
	description := []string{}
	var http_status int
	var status *model.StatusResponse

	if error != nil {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest
		status = &model.StatusResponse{
			HttpStatusCode: http.StatusBadRequest,
			ResponseCode:   general.ErrorStatusCode,
			Description:    description,
		}

	} else {
		_, error = controller.listService.UpdateList(request, context)

		if error == nil {

			description = append(description, "Success")
			http_status = http.StatusOK
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusOK,
				ResponseCode:   general.SuccessStatusCode,
				Description:    description,
			}

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest
			status = &model.StatusResponse{
				HttpStatusCode: http.StatusBadRequest,
				ResponseCode:   general.ErrorStatusCode,
				Description:    description,
			}
		}
	}

	return context.JSON(http_status, model.StandardResponse{
		Status: *status,
	})
}
