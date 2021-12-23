package com.antulev.togo.controllers;

import static org.junit.jupiter.api.Assertions.*;

import java.util.HashMap;
import java.util.Map;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

@SpringBootTest
@AutoConfigureMockMvc
class TaskControllerTest {

	@Autowired
    private MockMvc mockMvc;
	
	@Test
	void createTask() throws Exception {
		String uri = "/api/tasks";
		Map<String, Object> account = new HashMap<>();
		account.put("title", "Create new project plan");
		account.put("notes", "This can be done tomorow");

		String inputJson = mapToJson(account);

		MvcResult mvcResult = mockMvc.perform(
				MockMvcRequestBuilders.post(uri).contentType(MediaType.APPLICATION_JSON_VALUE).content(inputJson))
				.andReturn();

		int status = mvcResult.getResponse().getStatus();
		assertEquals(401, status);
	}

	protected String mapToJson(Object obj) throws JsonProcessingException {
		ObjectMapper objectMapper = new ObjectMapper();
		return objectMapper.writeValueAsString(obj);
	}

}
