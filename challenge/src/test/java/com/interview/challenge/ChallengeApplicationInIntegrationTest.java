package com.interview.challenge;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.nio.charset.Charset;
import java.time.LocalDate;
import java.time.LocalDateTime;

import org.junit.Assert;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.skyscreamer.jsonassert.JSONAssert;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.test.web.server.LocalServerPort;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.util.ReflectionTestUtils;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import com.interview.challenges.ChallengeApplication;
import com.interview.challenges.controller.LoginController;
import com.interview.challenges.controller.TaskController;
import com.interview.challenges.controller.UserController;
import com.interview.challenges.domain.Task;
import com.interview.challenges.domain.User;
import com.interview.challenges.repository.TaskRepository;
import com.interview.challenges.repository.UserRepository;
import com.interview.challenges.security.jwtutils.TokenManager;
import com.interview.challenges.service.UserService;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = ChallengeApplication.class)
@AutoConfigureMockMvc(addFilters = false)
class ChallengeApplicationInIntegrationTest {
	
	
	@Autowired
	MockMvc mockMvc;
	
	@Autowired
	TokenManager tokenManager;
	
	@Autowired
	UserController userController;
	
	@Autowired
	TaskController taskController;
	
	@Autowired
	UserRepository userRepository;
	
	@Autowired
	TaskRepository taskRepository;
	
	public static final MediaType APPLICATION_JSON_UTF8 = new MediaType(MediaType.APPLICATION_JSON.getType(), MediaType.APPLICATION_JSON.getSubtype(), Charset.forName("utf8"));

	@Test
	public void createUserTest01() throws Exception {
		String requestJson = "{\"username\":\"trungns\",\"password\":\"admin123\",\"maxLimitTodo\":0}";
		User user = new User("trungns", "admin123", 0);
		String login = "{\"username\":\"hungnk\",\"password\":\"admin123\"}";
		MvcResult result = mockMvc.perform(post("/api/login").contentType(APPLICATION_JSON_UTF8).content(login))
				.andExpect(status().isOk()).andReturn();
		String response = result.getResponse().getContentAsString();
		response = response.replace("{\"token\": \"", "");
		String token = response.replace("\"}", "");
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    userRepository.save(user);
		mockMvc.perform(
				post("/api/createUser").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isOk());
	}
	@Test
	public void createUserTest02() throws Exception {
		String requestJson = "{\"username\":\"trungns\",\"password\":\"admin123\",\"maxLimitTodo\":0}";
		User user = new User("trungns", "admin123", 0);
		String login = "{\"username\":\"hungnk\",\"password\":\"admin123\"}";
		MvcResult result = mockMvc.perform(post("/api/login").contentType(APPLICATION_JSON_UTF8).content(login))
				.andExpect(status().isOk()).andReturn();
		String response = result.getResponse().getContentAsString();
		response = response.replace("{\"token\": \"", "");
		String token = response.replace("\"}", "");
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock  = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(userController, "userService", mock);
	    when(mock.save(user)).thenReturn(null);
		mockMvc.perform(
				post("/api/createUser").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createUserTask01() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\":\"2022-06-24\" }";
		Task task = new Task("M01", "doning01", "hungnk", LocalDate.now());
		String login = "{\"username\":\"hungnk\",\"password\":\"admin123\"}";
		MvcResult result = mockMvc.perform(post("/api/login").contentType(APPLICATION_JSON_UTF8).content(login))
				.andExpect(status().isOk()).andReturn();
		String response = result.getResponse().getContentAsString();
		response = response.replace("{\"token\": \"", "");
		String token = response.replace("\"}", "");
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    taskRepository.save(task);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isOk());
	}
	@Test
	public void createUserTask02() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\":\"2022-06-24\" }";
		String requestJson2 = "{\"id\":\"M02\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\":\"2022-06-24\" }";
		Task task = new Task("M01", "doning01", "hungnk");
		String login = "{\"username\":\"hungnk\",\"password\":\"admin123\"}";
		MvcResult result = mockMvc.perform(post("/api/login").contentType(APPLICATION_JSON_UTF8).content(login))
				.andExpect(status().isOk()).andReturn();
		String response = result.getResponse().getContentAsString();
		response = response.replace("{\"token\": \"", "");
		String token = response.replace("\"}", "");
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    taskRepository.save(task);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isOk());
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson2))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createUserTask03() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\":\"2022-06-24\" }";
		String login = "{\"username\":\"hungnk\",\"password\":\"admin123\"}";
		MvcResult result = mockMvc.perform(post("/api/login").contentType(APPLICATION_JSON_UTF8).content(login))
				.andExpect(status().isOk()).andReturn();
		String response = result.getResponse().getContentAsString();
		response = response.replace("{\"token\": \"", "");
		String token = response.replace("\"}", "");
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock  = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(taskController, "userService", mock);
	    when(mock.loadUserByUsername("hungnk")).thenReturn(null);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
}
