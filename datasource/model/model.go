package model

import (
	"time"

	"github.com/AnanievNikolay/test_task/presentation/model"
	"gorm.io/gorm"
)

//New ...
func New(_response *model.ExternalAPIResponse) *DBModel {
	return &DBModel{
		APIResponse: _response.Marshal(),
		ReceivedAt:  time.Now().UTC(),
	}
}

//DBModel ...
type DBModel struct {
	APIResponse string    `gorm:"column:response"`
	ReceivedAt  time.Time `gorm:"column:received_at"`
}

//DBModelTable ...
func DBModelTable(dm *DBModel) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tableName := "test.currency_quotes"
		return tx.Table(tableName)
	}
}
