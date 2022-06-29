package com.api.todo.controllers;

import static org.hamcrest.CoreMatchers.hasItem;
//import static org.hamcrest.Matchers.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.test.mock.mockito.SpyBean;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockHttpServletRequestBuilder;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import com.api.todo.request.RequestTaskEntity;
import com.api.todo.services.TodoService;
import com.api.todo.services.UserService;
import com.api.todo.validators.TaskValidator;
import com.fasterxml.jackson.databind.ObjectMapper;

@RunWith(SpringRunner.class)
@WebMvcTest(TodoController.class)
public class TodoControllerTest {
	@Autowired
	MockMvc mockMvc;
	@Autowired
	ObjectMapper mapper;
	@MockBean
	TodoService todoService;
	@MockBean
	UserService userService;
	@SpyBean
	TaskValidator taskValidator;

	@Test
	public void createTask_failed_with_title_is_empty() throws Exception {
		RequestTaskEntity requestTask = new RequestTaskEntity("", "description_test", 123);

		MockHttpServletRequestBuilder mockRequest = MockMvcRequestBuilders.post("/api/todo/tasks")
				.contentType(MediaType.APPLICATION_JSON).accept(MediaType.APPLICATION_JSON)
				.content(this.mapper.writeValueAsString(requestTask));

		mockMvc.perform(mockRequest).andExpect(status().isBadRequest()).andExpect(jsonPath("$.errors").isArray())
				.andExpect(jsonPath("$.errors", hasItem("title field can't be empty.")));
	}

	@Test
	public void createTask_failed_with_description_is_empty() throws Exception {
		RequestTaskEntity requestTask = new RequestTaskEntity("title", "", 123);

		MockHttpServletRequestBuilder mockRequest = MockMvcRequestBuilders.post("/api/todo/tasks")
				.contentType(MediaType.APPLICATION_JSON).accept(MediaType.APPLICATION_JSON)
				.content(this.mapper.writeValueAsString(requestTask));

		mockMvc.perform(mockRequest).andExpect(status().isBadRequest()).andExpect(jsonPath("$.errors").isArray())
				.andExpect(jsonPath("$.errors", hasItem("description field can't be empty.")));
	}

	@Test
	public void createTask_failed_with_user_is_negative_number() throws Exception {
		RequestTaskEntity requestTask = new RequestTaskEntity("title", "fdsf", -123);

		MockHttpServletRequestBuilder mockRequest = MockMvcRequestBuilders.post("/api/todo/tasks")
				.contentType(MediaType.APPLICATION_JSON).accept(MediaType.APPLICATION_JSON)
				.content(this.mapper.writeValueAsString(requestTask));

		mockMvc.perform(mockRequest).andExpect(status().isBadRequest()).andExpect(jsonPath("$.errors").isArray())
				.andExpect(jsonPath("$.errors", hasItem("userId should not be empty or negative number.")));
	}

	@Test
	public void createTask_failed_with_user_is_not_existed() throws Exception {
		RequestTaskEntity requestTask = new RequestTaskEntity("title", "fdsf", 12338173872187382L);

		MockHttpServletRequestBuilder mockRequest = MockMvcRequestBuilders.post("/api/todo/tasks")
				.contentType(MediaType.APPLICATION_JSON).accept(MediaType.APPLICATION_JSON)
				.content(this.mapper.writeValueAsString(requestTask));

		mockMvc.perform(mockRequest).andExpect(status().isBadRequest()).andExpect(jsonPath("$.errors").isArray())
				.andExpect(jsonPath("$.errors", hasItem("userId field should be existed in table user.")));
	}

	@Test
     public void createTask_failed_with_title_is_maxlength_50() throws Exception {
     RequestTaskEntity requestTask = new RequestTaskEntity(
     "title title title titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle titletitle title",
     "description_test", 123);
    
     MockHttpServletRequestBuilder mockRequest = MockMvcRequestBuilders.post("/api/todo/tasks")
     .contentType(MediaType.APPLICATION_JSON).accept(MediaType.APPLICATION_JSON)
     .content(this.mapper.writeValueAsString(requestTask));
    
     mockMvc.perform(mockRequest).andExpect(status().isBadRequest()).andExpect(jsonPath("$.errors").isArray())
     .andExpect(jsonPath("$.errors", hasItem("title field has max length is 50.")));
     }

	@Test
	public void createTask_failed_with_description_is_maxlength_250() throws Exception {
		RequestTaskEntity requestTask = new RequestTaskEntity("title",
				"description_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_testdescription_test",
				123);

		MockHttpServletRequestBuilder mockRequest = MockMvcRequestBuilders.post("/api/todo/tasks")
				.contentType(MediaType.APPLICATION_JSON).accept(MediaType.APPLICATION_JSON)
				.content(this.mapper.writeValueAsString(requestTask));

		mockMvc.perform(mockRequest).andExpect(status().isBadRequest()).andExpect(jsonPath("$.errors").isArray())
				.andExpect(jsonPath("$.errors", hasItem("description field has max length is 250.")));
	}
}
