package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"svc-todo/entity"
	"svc-todo/model"
	"svc-todo/repository"
	"time"

	"github.com/labstack/echo"
)

type ListServiceInterface interface {
	CreateList(request *model.CreateListRequest, context echo.Context) (entity.List, error)
	GetList(request *model.GetListRequest) ([]entity.List, int, error)
	GetListWithSub(request *model.GetListRequest) ([]model.GetListResponse, int, error)
	UpdateList(request *model.UpdateListRequest, context echo.Context) (entity.List, error)
	DeleteList(request *int) error
}

type listService struct {
	listRepository repository.ListRepositoryInterface
}

func ListService(listRepository repository.ListRepositoryInterface) *listService {
	return &listService{listRepository}
}

func (service *listService) CreateList(request *model.CreateListRequest, context echo.Context) (entity.List, error) {
	var list entity.List
	var fileType string

	dir := os.Getenv("FILE_DIR")
	path := dir + "/list"
	file, error := context.FormFile("attachment")
	fileName := "-"

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error = os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if error == nil {
		src, err := file.Open()

		if err == nil {
			fileByte, _ := ioutil.ReadAll(src)
			fileType = http.DetectContentType(fileByte)
			split := strings.Split(fileType, "/")

			if fileType == "application/pdf" || fileType == "application/txt" {
				fileName = strconv.FormatInt(time.Now().Unix(), 10) + "." + split[1]
				error = ioutil.WriteFile(path+"/"+fileName, fileByte, 0777)
			} else {
				error = fmt.Errorf("Only pdf and txt can be allowed!")
			}
		}

		defer src.Close()
	}

	list = entity.List{
		Title:       request.Title,
		Description: request.Description,
		Attachment:  fileName,
	}
	if error == nil {
		_, error = service.listRepository.CreateList(&list)
	}

	return list, error
}

func (service *listService) GetList(request *model.GetListRequest) ([]entity.List, int, error) {
	var list []entity.List

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}

	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := service.listRepository.CountList()
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))
	url := os.Getenv("FILE_URL")

	if error == nil {
		list, error = service.listRepository.GetList(request)
		for index, value := range list {
			list[index].Attachment = url + "list/" + value.Attachment
		}
	}

	return list, int(total_pages), error
}

func (service *listService) GetListWithSub(request *model.GetListRequest) ([]model.GetListResponse, int, error) {
	var list []entity.List
	var result []model.GetListResponse

	if request.PageSize == 0 {
		request.PageSize = math.MaxInt16
	}

	request.StartIndex = request.PageNo * request.PageSize
	total_data, error := service.listRepository.CountList()
	total_pages := math.Ceil(float64(total_data) / float64(request.PageSize))
	url := os.Getenv("FILE_URL")

	if error == nil {
		list, error = service.listRepository.GetListWithSub(request)

		for _, value := range list {
			var sub_list []*entity.SubList
			json.Unmarshal([]byte(value.SubList), &sub_list)

			result = append(result, model.GetListResponse{Id: value.Id, Title: value.Title, Description: value.Description, Attachment: url + "list/" + value.Attachment, SubList: sub_list})
		}
	}

	return result, int(total_pages), error
}

func (service *listService) UpdateList(request *model.UpdateListRequest, context echo.Context) (entity.List, error) {
	var list entity.List
	var fileType string

	dir := os.Getenv("FILE_DIR")
	path := dir + "/list"
	file, error := context.FormFile("attachment")
	fileName := "-"

	_, check_dir_error := os.Stat(path)

	if os.IsNotExist(check_dir_error) {
		check_dir_error = os.MkdirAll(path, 0755)

		if check_dir_error != nil {
			error = check_dir_error
		}
	}

	if error == nil {
		src, err := file.Open()

		if err == nil {
			fileByte, _ := ioutil.ReadAll(src)
			fileType = http.DetectContentType(fileByte)
			split := strings.Split(fileType, "/")

			if fileType == "application/pdf" || fileType == "application/txt" {
				fileName = strconv.FormatInt(time.Now().Unix(), 10) + "." + split[1]
				error = ioutil.WriteFile(path+"/"+fileName, fileByte, 0777)
			} else {
				error = fmt.Errorf("Only pdf and txt can be allowed!")
			}
		}

		defer src.Close()
	}

	list = entity.List{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		Attachment:  fileName,
	}
	if error == nil {
		_, error = service.listRepository.UpdateList(&list)
	}

	return list, error
}

func (service *listService) DeleteList(request *int) error {

	error := service.listRepository.DeleteList(request)

	return error
}
