package transformer

import (
	"errors"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"reflect"
)

type Resource struct {
	Transformer Transformer
}

type Resourcer interface {
	GetTransformer() any
	SetTransformer(transformer Transformer)
}

func (r *Resource) GetTransformer() any {
	return r.Transformer
}

func (r *Resource) SetTransformer(transformer Transformer) {
	r.Transformer = transformer
}

type Item struct {
	Resource
	Data any
}

type Collection struct {
	Resource
	Total        int
	Page         int
	ItemsPerPage int
	Data         []any
}

type Error struct {
	Resource
	Data any
}

func (i *Item) GetData() any {
	return i.Data
}

func (i *Item) SetData(data any) {
	i.Data = data
}

func (c *Collection) GetData() []any {
	return c.Data
}

func (c *Collection) SetData(data []any) {
	c.Data = data
}

func (c *Collection) GetTotal() int {
	return c.Total
}

func (c *Collection) SetTotal(total int) {
	c.Total = total
}

func (c *Collection) GetPage() int {
	return c.Page
}

func (c *Collection) SetPage(page int) {
	c.Page = page
}

func (c *Collection) GetItemsPerPage() int {
	return c.ItemsPerPage
}

func (c *Collection) SetItemsPerPage(itemsPerPage int) {
	c.ItemsPerPage = itemsPerPage
}

func (e *Error) GetData() any {
	return e.Data
}

func (e *Error) SetData(data any) {
	e.Data = data
}

// NewItem /**
func NewItem(data any, transformer Transformer) Item {
	i := Item{}
	i.SetData(data)
	i.SetTransformer(transformer)
	return i
}

// NewCollection /**
func NewCollection(data any, transformer Transformer) Collection {
	c := Collection{}
	var r = make([]any, 0)
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			e := s.Index(i).Interface()
			r = append(r, e)
		}
	default:
		utils.PanicInternalServerError(errors.New("can not reflect value type"))
	}
	c.SetData(r)
	c.SetTransformer(transformer)
	return c
}

// NewError /**
func NewError(data any) Error {
	e := Error{}
	e.SetData(data)
	e.SetTransformer(nil)
	return e
}
