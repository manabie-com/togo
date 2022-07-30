package com.uuhnaut69.app.todo.api;

import static org.hamcrest.Matchers.is;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.uuhnaut69.app.todo.model.Todo;
import com.uuhnaut69.app.todo.model.dto.TodoRequest;
import com.uuhnaut69.app.todo.service.TodoService;
import java.time.Instant;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

/**
 * @author uuhnaut
 */
@RunWith(SpringRunner.class)
@WebMvcTest(TodoController.class)
class TodoControllerTest {

  @Autowired
  private MockMvc mvc;

  @MockBean
  private TodoService todoService;

  @Test
  void createNewTodo() throws Exception {
    var todoRequest = new TodoRequest("Test", 1L);
    var todo = new Todo(1L, "Test", 1L, Instant.now());

    Mockito.when(todoService.createNewTodo(todoRequest)).thenReturn(todo);

    mvc.perform(MockMvcRequestBuilders.post("/todos")
            .contentType(MediaType.APPLICATION_JSON_VALUE)
            .content(new ObjectMapper().writeValueAsString(todoRequest)))
        .andExpect(MockMvcResultMatchers.status().isCreated())
        .andExpect(MockMvcResultMatchers.jsonPath("$.id", is(1)))
        .andExpect(MockMvcResultMatchers.jsonPath("$.task", is("Test")));
  }

  @Test
  void createNewTodoFailedByTaskIsBlank() throws Exception {
    var todoRequest = new TodoRequest("", 1L);

    mvc.perform(MockMvcRequestBuilders.post("/todos")
            .contentType(MediaType.APPLICATION_JSON_VALUE)
            .content(new ObjectMapper().writeValueAsString(todoRequest)))
        .andExpect(MockMvcResultMatchers.status().isBadRequest());
  }

  @Test
  void createNewTodoFailedByUserIdIsNull() throws Exception {
    var todoRequest = new TodoRequest("Test", null);

    mvc.perform(MockMvcRequestBuilders.post("/todos")
            .contentType(MediaType.APPLICATION_JSON_VALUE)
            .content(new ObjectMapper().writeValueAsString(todoRequest)))
        .andExpect(MockMvcResultMatchers.status().isBadRequest());
  }

  @Test
  void createNewTodoFailedByTaskIsBlankAndUserIdIsNull() throws Exception {
    var todoRequest = new TodoRequest("", null);

    mvc.perform(MockMvcRequestBuilders.post("/todos")
            .contentType(MediaType.APPLICATION_JSON_VALUE)
            .content(new ObjectMapper().writeValueAsString(todoRequest)))
        .andExpect(MockMvcResultMatchers.status().isBadRequest());
  }
}