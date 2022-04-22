package utils

import (
	"crud_v2/model"
	"math"

	"github.com/jinzhu/gorm"
)

func Paginate(value interface{}, pagination *model.Pagination, db *gorm.DB, preload string, selectDataPreload string, whtsSelect ...string) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Select("id").Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		if preload != "" {
			db = db.Preload(preload)
		}
		if preload != "" && selectDataPreload != "" {
			db = db.Preload(preload, func(db *gorm.DB) *gorm.DB {
				return db.Select(selectDataPreload)
			})
		}

		if whtsSelect != nil {
			db = db.Select(whtsSelect)
		}

		return db.Model(value).Limit(pagination.GetLimit()).Offset(pagination.GetOffset())
	}
}
