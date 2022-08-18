package domain

type WaiterUpsertrequest struct {
	UUID		string	`json:"uuid"`
	Name		string	`json:"name"`
	CreatedAt	string	`json:"created_at"`
}

type WaiterResponse struct {
	BranchUUID		string	`json:"branch_uuid"`
	Name			string	`json:"name"`
	LastActive		*int64	`json:"last_active,omitempty"`
	CreatedAt		int64	`json:"created_at"`
	UpdatedAt		*int64	`json:"updated_at,omitempty"`
	DeletedAt		*int64	`json:"deleted_at,omitempty"`
}