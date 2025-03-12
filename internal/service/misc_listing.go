package service

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScopeFunc = func(*gorm.DB) *gorm.DB

type Pagination struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (p *Pagination) String() string {
	return fmt.Sprintf("Pagination(page=%d, page_size=%d)", p.Page, p.PageSize)
}

func (p *Pagination) WithDefault() *Pagination {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 || p.PageSize >= 100 {
		p.PageSize = 10
	}

	return p
}

func (p *Pagination) Scope() ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.PageSize * (p.Page - 1)).Limit(p.PageSize)
	}
}

type OrderBy struct {
	Column     string `json:"column" form:"column"`
	Descending bool   `json:"descending" form:"descending"`
}

func (o *OrderBy) String() string {
	return fmt.Sprintf("OrderBy(column=%s, descending=%t", o.Column, o.Descending)
}

func (o *OrderBy) Scope() ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		var column clause.Column
		if o.Column == "" {
			column = clause.PrimaryColumn
		} else {
			column = clause.Column{Name: o.Column}
		}

		return db.Order(
			clause.OrderBy{
				Columns: []clause.OrderByColumn{
					{Column: column, Desc: o.Descending},
				},
			},
		)
	}
}
