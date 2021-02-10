package unitofwork

import (
	"container/list"
	"log"
	"sync"

	"github.com/AnanievNikolay/test_task/datasource/repository"
)

//UnitOfWork ...
type UnitOfWork struct {
	dirty      *list.List
	repository repository.IRepository
	mutex      *sync.Mutex
}

//NewDirty ...
func (uow *UnitOfWork) NewDirty(_dirty interface{}) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.dirty.PushBack(_dirty)
}

//Commit ...
func (uow *UnitOfWork) Commit() {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	if !uow.repository.Save() {
		log.Println("[UnitOfWork] Something went wrong wile commit")
		return
	}
	uow.dirty = list.New()
}
