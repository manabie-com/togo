package todorepo

import (
	"errors"
	"log"

	"github.com/hashicorp/go-memdb"
	"github.com/manabie-com/togo/api/model"
	"github.com/manabie-com/togo/internal/repo"
)

type InmemTodo struct {
	Conn repo.Conn
}

func GetTodoSchema() *memdb.DBSchema {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"todo": {
				Name: "todo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:         "id",
						AllowMissing: false,
						Unique:       true,
						Indexer:      &memdb.StringFieldIndex{Field: "ID"},
					},
					"date": {
						Name:         "date",
						AllowMissing: false,
						Unique:       false,
						Indexer:      &memdb.StringFieldIndex{Field: "CreatedDate"},
					},
					"user": {
						Name:         "user",
						AllowMissing: false,
						Unique:       false,
						Indexer:      &memdb.StringFieldIndex{Field: "User"},
					},
				},
			},
			"user": {
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:         "id",
						Unique:       true,
						AllowMissing: false,
						Indexer:      &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
	err := schema.Validate()
	if err != nil {
		panic(err)
	}
	return schema
}

func (conn *InmemTodo) Add(t model.Todo) (model.Todo, error) {
	tx, err := conn.Conn.GetTxn()
	if err != nil {
		log.Print(err)
		return model.Todo{}, errors.New("unable to get stroage transaction")
	}
	if err := tx.(*memdb.Txn).Insert("todo", t); err != nil {
		log.Print(err)
		return model.Todo{}, errors.New("unable to add todo on storage")
	}
	tx.(*memdb.Txn).Commit()
	return t, nil
}

func (conn *InmemTodo) Update(t model.Todo) (int, error) {
	return 0, nil
}
func (conn *InmemTodo) Delete(ID string) error {
	return nil
}

func (conn *InmemTodo) GetOne(ID string) (model.Todo, error) {
	return model.Todo{}, nil
}

func (conn *InmemTodo) GetByUserAndDate(ID, date string) ([]model.Todo, error) {
	return []model.Todo{}, nil
}

func (conn *InmemTodo) Get(ids []string) ([]model.Todo, error) {
	tx, err := conn.Conn.GetTxn()
	if err != nil {
		log.Print(err)
		return []model.Todo{}, errors.New("unable to get stroage transaction")
	}
	r, err := tx.(*memdb.Txn).Get("todo", "id")
	if err != nil {
		log.Print(err)
		return []model.Todo{}, errors.New("unable to get todos")
	}
	defer tx.(*memdb.Txn).Abort()
	res := []model.Todo{}
	for obj := r.Next(); obj != nil; obj = r.Next() {
		res = append(res, obj.(model.Todo))
	}
	return res, nil
}
