package domain

type FileableType string

const (
	FileableType_ITEM FileableType = "item_categories.items"
)

type UploadFileRequest struct {
	UUID         string       `json:"uuid"`
	FileableUUID string       `json:"fileable_uuid"`
	FileableType FileableType `json:"fileable_type"`
}

type UploadFileResponse struct {
	UUID         string       `json:"uuid"`
	BranchUUID   string       `json:"branch_uuid"`
	FileableUUID string       `json:"fileable_uuid"`
	FileableType FileableType `json:"fileable_type"`
	Url          string       `json:"url"`
	Path         string       `json:"path"`
	Ext          string       `json:"ext"`
}
