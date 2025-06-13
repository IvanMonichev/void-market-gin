package transport

type Meta struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type PaginatedResponse[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}
