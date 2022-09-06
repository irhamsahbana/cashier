package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func (u *FileUploadUsecase) UploadFile(c context.Context, branchId string, file *multipart.FileHeader, req *domain.UploadFileRequest) (*domain.UploadFileResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.FileableUUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	if err := validateFileableType(req.FileableType); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data domain.File
	data.UUID = req.UUID
	data.BranchUUID = branchId
	data.FileableUUID = req.FileableUUID
	data.FileableType = req.FileableType
	data.Ext = filepath.Ext(file.Filename)
	data.Path = "storage/" + string(data.FileableType) + "/" + data.UUID + data.Ext

	result, code, err := u.fileRepo.UploadFile(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.UploadFileResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.FileableUUID = result.FileableUUID
	resp.FileableType = result.FileableType
	resp.Ext = result.Ext
	resp.Path = result.Path

	// save file *multipart.FileHeader to storage

	return &resp, http.StatusOK, nil
}

func validateFileableType(ft domain.FileableType) error {
	switch ft {
	case "item_categories.items":
		return nil
	default:
		return errors.New("Invalid fileable type")
	}
}
