package repository

import "github.com/AnanievNikolay/test_task/presentation/model"

//IRepository ...
type IRepository interface {
	Save(*model.ExternalAPIResponse) bool
}
