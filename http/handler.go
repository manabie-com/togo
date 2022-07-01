package http

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/lawtrann/togo"
)

type Handler struct {
	TodoService togo.TodoService
	TodoPattern *regexp.Regexp
}

func NewHandler(todoService togo.TodoService) *Handler {
	return &Handler{
		TodoService: todoService,
		TodoPattern: regexp.MustCompile(`^/api/(?P<user_name>\w+)/todos/?$`),
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.TodoPattern.MatchString(r.URL.Path) {
		switch r.Method {
		case http.MethodPost:
			matches := h.TodoPattern.FindStringSubmatch(r.URL.Path)
			result := make(map[string]string)
			for i, name := range h.TodoPattern.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = matches[i]
				}
			}
			userName := strings.ToLower(result["user_name"])

			h.CreateTodo(userName, w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) parserRequest(r *http.Request) (togo.Todo, error) {
	dec := json.NewDecoder(r.Body)
	var t togo.Todo
	err := dec.Decode(&t)
	if err != nil {
		return togo.Todo{}, err
	}
	return t, nil
}

func (h *Handler) CreateTodo(userName string, w http.ResponseWriter, r *http.Request) {
	t, err := h.parserRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Todo object"))
		return
	}

	// Check if empty description
	if len(t.Description) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty Todo Description!"))
		return
	}

	res, err := h.TodoService.AddTodoByUser(userName, &t)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(res, w)
}
