package uow

import (
	"context"
	"database/sql"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uwo *UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	tx, _ := db.BeginTx(ctx, nil)
	return &Uow{
		Db:           db,
		Tx:           tx,
		Repositories: make(map[string]RepositoryFactory),
	}

}

func (uow *Uow) Register(name string, fc RepositoryFactory) {
	uow.Repositories[name] = fc
}

func (uow *Uow) UnRegister(name string) {
	delete(uow.Repositories, name)
}
