package unitofwork

import (
	"container/list"
	"log"
	"sync"

	"github.com/AnanievNikolay/test_task/datasource/repository"
	"github.com/AnanievNikolay/test_task/presentation/model"
)

//New ..
func New(_repository repository.IRepository) IUnitOfWork {
	return &UnitOfWork{
		dirty:      list.New(),
		repository: _repository,
		mutex:      &sync.Mutex{},
	}
}

//UnitOfWork ...
type UnitOfWork struct {
	dirty      *list.List
	repository repository.IRepository
	mutex      *sync.Mutex
}

//NewDirty ...
func (uow *UnitOfWork) NewDirty(_dirty *model.ExternalAPIResponse) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.dirty.PushBack(_dirty)
}

//Commit ...
func (uow *UnitOfWork) Commit() {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	for el := uow.dirty.Front(); el != nil; el = el.Next() {
		if uow.repository.Save(el.Value.(*model.ExternalAPIResponse)) {
			uow.dirty.Remove(el)
		} else {
			log.Println("[Error] Something went wrong while commit")
			return
		}
	}
}
