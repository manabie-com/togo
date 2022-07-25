package transformer

import "math"

type Serializer interface {
	SerializeItem(resource Item) any
	SerializeCollection(resource Collection) any
	SerializeError(resource Error) any
	SerializeNil() any
}

// NewJSONSerializer /**
func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{
		Status:   1,
		Messages: []string{"Successful."},
		Data:     nil,
	}
}

type JSONSerializer struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
	Data     any      `json:"data"`
}

type CollectionJson struct {
	Items      []any          `json:"items"`
	Pagination PaginationJson `json:"-"`
}

type PaginationJson struct {
	Current      int `json:"current"`
	Prev         int `json:"prev"`
	Next         int `json:"next"`
	TotalItems   int `json:"total_items"`
	First        int `json:"first"`
	Last         int `json:"last"`
	ItemsPerPage int `json:"items_per_page"`
}

func (j JSONSerializer) SerializeItem(resource Item) interface{} {
	formattedItem := resource.Transformer.Transform(resource.Data)
	j.Data = formattedItem

	return j
}

func (j JSONSerializer) SerializeCollection(resource Collection) interface{} {
	var formattedCollection = make([]any, 0)
	for _, v := range resource.GetData() {
		formattedCollection = append(formattedCollection, resource.Transformer.Transform(v))
	}

	collection := CollectionJson{}
	total := resource.GetTotal()               // total
	page := resource.GetPage()                 // page from request
	itemsPerPage := resource.GetItemsPerPage() // items_per_page from request
	last := 1                                  // last default is 1
	if itemsPerPage != 0 {
		last = int(math.Ceil(float64(total / itemsPerPage))) // calculate the last
	} else {
		itemsPerPage = total
	}
	collection.Items = formattedCollection
	collection.Pagination = PaginationJson{
		Current:      page,
		Prev:         1,
		Next:         1,
		First:        1,
		Last:         last,
		TotalItems:   total,
		ItemsPerPage: itemsPerPage,
	}

	j.Data = collection

	return j
}

func (j JSONSerializer) SerializeNil() any {
	return j
}

func (j JSONSerializer) SerializeError(resource Error) any {
	payload := resource.GetData()
	j.Status = 0 // default 0 for all error
	j.Data = payload
	return j
}
