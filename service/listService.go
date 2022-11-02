package service

import (
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

	date_now := time.Now()
	dir := os.Getenv("FILE_DIR")
	path := dir + "/list/" + date_now.Format("2006-01-02")
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

	if error == nil {
		list, error = service.listRepository.GetList(request)
	}

	return list, int(total_pages), error
}
