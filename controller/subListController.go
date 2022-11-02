package controller

import (
	"net/http"
	"svc-todo/general"
	"svc-todo/model"
	"svc-todo/service"

	"github.com/labstack/echo"
)

type subListController struct {
	subListService service.SubListServiceInterface
}

func SubListController(subListService service.SubListServiceInterface) *subListController {
	return &subListController{subListService}
}

func (controller *subListController) CreateSubList(context echo.Context) error {

	request := new(model.CreateSubListRequest)

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
		_, error = controller.subListService.CreateSubList(request, context)

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
