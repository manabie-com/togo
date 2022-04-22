package helper

// Problem data response
type Problem struct {
	Status  int         `json:"status"`
	Title   string      `json:"title"`
	Details interface{} `json:"details"`
}

// Success data response
type Success struct {
	Status int         `json:"status"`
	Mess   string      `json:"Mess"`
	Data   interface{} `json:"data,omitempty"`
}
