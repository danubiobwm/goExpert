package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uwo *Uow) error) error
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

func (uow *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if uow.Tx == nil {
		tx, err := uow.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		uow.Tx = tx
	}
	repo := uow.Repositories[name](uow.Tx)
	return repo, nil
}

func (uow *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {

	if uow.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := uow.Db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	uow.Tx = tx

	err = fn(uow)

	if err != nil {
		errRb := uow.Rollback()

		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}

		return err
	}
	return uow.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.Tx != nil {
		return errors.New("No transaction to rollback")
	}
	err := u.Tx.Rollback()
	if err != nil {
		return err
	}
	u.Tx = nil

	return nil
}

func (uow *Uow) CommitOrRollback() error {
	err := uow.Tx.Commit()

	if err != nil {
		errRb := uow.Rollback()

		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}

		return err
	}
	uow.Tx = nil
	return nil
}
