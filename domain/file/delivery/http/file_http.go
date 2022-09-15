package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type FileUploadHandler struct {
	FileUsecase   domain.FileUsecaseContract
	appStorageURL string
}

func NewFileHandler(router *gin.Engine, usecase domain.FileUsecaseContract, appStorageURL string) {
	handler := &FileUploadHandler{
		FileUsecase:   usecase,
		appStorageURL: appStorageURL,
	}

	permitted := []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
	}

	router.PUT("files", middleware.Auth, middleware.Authorization(permitted), handler.UploadFile)
}

func (h *FileUploadHandler) UploadFile(c *gin.Context) {
	var request domain.UploadFileRequest

	request.UUID = c.Request.FormValue("uuid")
	request.FileableUUID = c.Request.FormValue("fileable_uuid")
	request.FileableType = domain.FileableType(c.Request.FormValue("fileable_type"))
	branchId := c.GetString("branch_uuid")

	file, err := c.FormFile("file")
	if err != nil {
		http_response.ReturnResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.FileUsecase.UploadFile(ctx, branchId, file, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	filename := result.UUID + result.Ext
	relativePath := "./storage-uploads" + "/" + string(result.FileableType)
	dst := path.Join(relativePath, filename)

	if exists, err := exists(relativePath); !exists {
		if err != nil {
			http_response.ReturnResponse(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		if err := os.MkdirAll(relativePath, os.ModePerm); err != nil {
			http_response.ReturnResponse(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		http_response.ReturnResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	result.Url = h.appStorageURL + "/" + string(result.FileableType) + "/" + filename
	http_response.ReturnResponse(c, httpCode, "File uploaded successfully", result)
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
