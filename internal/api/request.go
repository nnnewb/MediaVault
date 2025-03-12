package api

import "github.com/nnnewb/media-vault/internal/service"

const (
	DefaultPage            = 1
	DefaultPageSize        = 10
	DefaultOrderColumn     = ""
	DefaultOrderDescending = true
	PageSizeMaximum        = 1000
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (p *Pagination) WithDefault() *Pagination {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.PageSize <= 0 || p.PageSize >= PageSizeMaximum {
		p.PageSize = DefaultPageSize
	}
	return p
}

func (p *Pagination) QueryOption() service.QueryOption {
	return service.WithPaginate(p.Page, p.PageSize)
}

type OrderBy struct {
	Column     string `json:"column"`
	Descending *bool  `json:"descending"`
}

func (o *OrderBy) WithDefault() *OrderBy {
	if o.Column == "" {
		o.Column = DefaultOrderColumn
	}

	if o.Descending == nil {
		o.Descending = new(bool)
		*o.Descending = DefaultOrderDescending
	}

	return o
}

func (o *OrderBy) QueryOption() service.QueryOption {
	descending := false
	if o.Descending != nil {
		descending = *o.Descending
	}
	return service.WithOrderBy(o.Column, descending)
}
