package unitofwork

import "github.com/AnanievNikolay/test_task/presentation/model"

type IUnitOfWork interface {
	NewDirty(*model.ExternalAPIResponse)
	Commit()
}
