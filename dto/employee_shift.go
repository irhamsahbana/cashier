package dto

type EntryCashInsertRequest struct {
	Username    string  `json:"username"`
	Description string  `json:"description"`
	Expense     bool    `json:"expense"`
	Value       float64 `json:"value"`
	CreatedAt   string  `json:"created_at"`
}
