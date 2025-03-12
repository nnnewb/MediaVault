package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QueryOption func(tx *gorm.DB)

func WithOrderBy(column string, desc bool) QueryOption {
	return func(tx *gorm.DB) {
		var orderByColumnClause clause.OrderByColumn
		if column != "" {
			orderByColumnClause = clause.OrderByColumn{
				Column: clause.Column{
					Name: column,
				},
				Desc: desc,
			}
		} else {
			orderByColumnClause = clause.OrderByColumn{
				Column: clause.PrimaryColumn,
				Desc:   desc,
			}
		}

		tx.Order(orderByColumnClause)
	}
}

func WithPaginate(page, pageSize int) QueryOption {
	return func(tx *gorm.DB) {
		if pageSize == 0 {
			pageSize = 10
		}

		if page == 0 {
			page = 1
		}

		tx.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
