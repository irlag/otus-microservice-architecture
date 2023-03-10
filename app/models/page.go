package models

import (
	"math"
	"net/http"
	"strconv"
)

const DefaultLimit int64 = 10
const DefaultPage int64 = 1

type Paginated interface {
	GetOffset() int64
	GetPage() int64
	GetLimit() int64
	GetPages() int64
	GetTotal() int64
	SetTotal(total int64)
}

type paginated struct {
	Limit int64
	Page  int64
	Total int64
}

func NewPaginated(limit int64, page int64) Paginated {
	if page <= 0 {
		page = DefaultPage
	}

	if limit <= 0 {
		limit = DefaultLimit
	}

	return &paginated{
		Limit: limit,
		Page:  page,
	}
}

func NewPaginatedFromRequest(request *http.Request) Paginated {
	queryParams := request.URL.Query()

	limit, _ := strconv.Atoi(queryParams.Get("limit"))
	page, _ := strconv.Atoi(queryParams.Get("page"))

	return NewPaginated(int64(limit), int64(page))
}

func (p *paginated) GetOffset() int64 {
	return (p.Page - 1) * p.Limit
}

func (p *paginated) GetPage() int64 {
	return p.Page
}

func (p *paginated) GetLimit() int64 {
	return p.Limit
}

func (p *paginated) GetPages() int64 {
	return int64(math.Ceil(float64(p.Total) / float64(p.Limit)))
}

func (p *paginated) GetTotal() int64 {
	return p.Total
}

func (p *paginated) SetTotal(total int64) {
	p.Total = total
}
