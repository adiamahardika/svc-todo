package service

import (
	"fmt"
	"io/ioutil"
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

type SubListServiceInterface interface {
	CreateSubList(request *model.CreateSubListRequest, context echo.Context) (entity.SubList, error)
	GetSubList(request *entity.SubList) ([]entity.SubList, error)
	UpdateSubList(request *model.UpdateSubListRequest, context echo.Context) (entity.SubList, error)
	DeleteSubList(request *int) error
}

type subListService struct {
	subListRepository repository.SubListRepositoryInterface
}

func SubListService(subListRepository repository.SubListRepositoryInterface) *subListService {
	return &subListService{subListRepository}
}

func (service *subListService) CreateSubList(request *model.CreateSubListRequest, context echo.Context) (entity.SubList, error) {

	var sub_list entity.SubList
	var fileType string

	dir := os.Getenv("FILE_DIR")
	path := dir + "/sub_list"
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

	sub_list = entity.SubList{
		Title:       request.Title,
		Description: request.Description,
		Attachment:  fileName,
		IdList:      request.IdList,
	}
	if error == nil {
		_, error = service.subListRepository.CreateSubList(&sub_list)
	}

	return sub_list, error
}

func (service *subListService) GetSubList(request *entity.SubList) ([]entity.SubList, error) {

	url := os.Getenv("FILE_URL")
	sub_list, error := service.subListRepository.GetSubList(request)

	if error == nil {
		for index, value := range sub_list {
			sub_list[index].Attachment = url + "sub_list/" + value.Attachment
		}
	}
	return sub_list, error
}

func (service *subListService) UpdateSubList(request *model.UpdateSubListRequest, context echo.Context) (entity.SubList, error) {

	var sub_list entity.SubList
	var fileType string

	dir := os.Getenv("FILE_DIR")
	path := dir + "/sub_list"
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

	sub_list = entity.SubList{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		Attachment:  fileName,
		IdList:      request.IdList,
	}
	if error == nil {
		_, error = service.subListRepository.UpdateSubList(&sub_list)
	}

	return sub_list, error
}

func (service *subListService) DeleteSubList(request *int) error {

	error := service.subListRepository.DeleteSubList(request)

	return error
}
