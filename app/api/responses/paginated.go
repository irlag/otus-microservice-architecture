package responses

type Paginated struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
	Pages int64 `json:"pages"`
}
