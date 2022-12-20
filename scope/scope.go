package scope

import (
	"fmt"

	"github.com/fenghaojiang/gorm-scope-utils/value"
	"gorm.io/gorm"
)

func ScopeRange[T comparable](rangeValues ...value.ValueRange[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var empty T
		for i := range rangeValues {
			if rangeValues[i].IncludeEmpty {
				db = db.Where(fmt.Sprintf("`%s` >= ? and `%s` < ?", rangeValues[i].Field, rangeValues[i].Field), rangeValues[i].From, rangeValues[i].To)
				continue
			}
			if rangeValues[i].From != empty {
				db = db.Where(fmt.Sprintf("`%s` >= ?", rangeValues[i].Field), rangeValues[i].From)
			}
			if rangeValues[i].To != empty {
				db = db.Where(fmt.Sprintf("`%s` < ?", rangeValues[i].Field), rangeValues[i].To)
			}
		}
		return db
	}
}

func ScopeIn[T comparable](inValues ...value.ValueIn[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for i := range inValues {
			if len(inValues[i].Values) == 0 {
				continue
			}
			db = db.Where(fmt.Sprintf("`%s` in ?", inValues[i].Field), inValues[i].Values)
		}
		return db
	}
}

func ScopeEqual[T comparable](equalValues ...value.ValueEqual[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var empty T
		for i := range equalValues {
			if equalValues[i].IncludeEmpty {
				db = db.Where(fmt.Sprintf("`%s` = ?", equalValues[i].Field), equalValues[i].Value)
				continue
			}
			if equalValues[i].Value != empty {
				db = db.Where(fmt.Sprintf("`%s` = ?", equalValues[i].Field), equalValues[i].Value)
			}
		}
		return db
	}
}
