package dto

type PaginatedResponse[T any] struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	Data       []T   `json:"data"`
}

func ToPaginatedList[T any](items []T, page, limit int, total int64) PaginatedResponse[T] {
	totalPages := int((total + int64(limit) - 1) / int64(limit)) // ceiling division

	return PaginatedResponse[T]{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Data:       items,
	}
}
