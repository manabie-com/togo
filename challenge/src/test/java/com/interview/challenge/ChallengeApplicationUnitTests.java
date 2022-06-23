package com.interview.challenge;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNotNull;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.nio.charset.Charset;
import java.time.LocalDate;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

import org.junit.Assert;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.util.ReflectionTestUtils;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;

import com.interview.challenges.ChallengeApplication;
import com.interview.challenges.controller.TaskController;
import com.interview.challenges.controller.UserController;
import com.interview.challenges.domain.Task;
import com.interview.challenges.domain.User;
import com.interview.challenges.repository.TaskRepository;
import com.interview.challenges.repository.UserRepository;
import com.interview.challenges.security.jwtutils.TokenManager;
import com.interview.challenges.service.TaskService;
import com.interview.challenges.service.UserService;
import com.interview.challenges.utils.MessageBody;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = ChallengeApplication.class)
@AutoConfigureMockMvc
class ChallengeApplicationUnitTests {
	
	@Autowired
	private UserService userService;
	
	@Autowired
	private UserRepository userRepository;
	
	@Autowired
	private TaskService taskService;
	
	@Mock
	private UserService userServiceMock;
	
	@Autowired
	private TokenManager tokenManager;
	
	@Autowired
	private UserController userController;
	
	@Autowired
	private TaskController taskController;
	
	
	@Autowired
	private TaskRepository taskRepository;
	
	@Autowired
	MockMvc mockMvc;
	
	public static final MediaType APPLICATION_JSON_UTF8 = new MediaType(MediaType.APPLICATION_JSON.getType(), MediaType.APPLICATION_JSON.getSubtype(), Charset.forName("utf8"));

	@Test
	public void createUserTest01() {
		User userBefore = new User("hungnk", "admin123", 0);
		User user = userService.save(userBefore);
		Assert.assertEquals(userBefore, user);
	}
	
	@Test
	public void createUserTest02() {
		User userBefore = new User("hungnk", "admin123", 0);
		userService.save(userBefore);
		User user = (User) userService.loadUserByUsername("hungnk");
		assertFalse(Objects.isNull(user));
        assertNotNull(user);
	}
	@Test
	public void createUserTest03() {
		User userBefore = new User("hungnk", "admin123", 0);
		userService.save(userBefore);
		try {
		User user = (User) userService.loadUserByUsername("trung");
		}catch (UsernameNotFoundException e) {
			Assert.assertEquals("User \'trung\' not found", e.getMessage());
		}
	}
	
	@Test
	public void createUserTest04() {
		User userBefore = new User("hungnk", "admin123", 0);
		User user = userRepository.save(userBefore);
		Assert.assertEquals(userBefore, user);
	}
	
	@Test
	public void createUserTest05() {
		User userBefore = new User("hungnk", "admin123", 0);
		userRepository.save(userBefore);
		User user = (User) userRepository.findByUsername("hungnk");
		assertFalse(Objects.isNull(user));
        assertNotNull(user);
	}
	
	@Test
	public void createUserTest06() throws Exception {
		String requestJson = "{\"username\":\"trungns\",\"password\":\"admin123\",\"maxLimitTodo\":0}";
		User userBefore = new User("hungnk", "admin123", 0);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock  = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(userController, "userService", mock);
        User user = new User("trungns", "admin123", 0);
	    when(mock.save(user)).thenReturn(user);
		mockMvc.perform(
				post("/api/createUser").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isOk());
	}
	
	@Test
	public void createUserTest07() throws Exception {
		String requestJson = "{\"username\":\"trungns\",\"password\":\"admin123\",\"maxLimitTodo\":0}";
		User userBefore = new User("hungnk", "admin123", 0);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock  = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(userController, "userService", mock);
        User user = new User("trungns", "admin123", 0);
	    when(mock.save(user)).thenReturn(null);
		mockMvc.perform(
				post("/api/createUser").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createTaskTest01() {
		User userBefore = new User("hungnk", "admin123", 0);
		userService.save(userBefore);
		Task beforTask = new Task("01", "doing", "hungnk");
		Task task = taskService.save(beforTask);
		Assert.assertEquals(beforTask, task);
	}
	
	@Test
	public void createTaskTest02() {
		User userBefore = new User("hungnk", "admin123", 0);
		userService.save(userBefore);
		Task beforTask = new Task("01", "doing", "hungnk");
		taskService.save(beforTask);
		List<Task> task = taskService.findByUserId("hungnk");
		assertFalse(task.isEmpty());
		assertNotNull(task);
	}
	
	@Test
	public void createTaskTest03() {
		User userBefore = new User("hungnk", "admin123", 01);
		userRepository.save(userBefore);
		Task beforTask = new Task("01", "doing", "hungnk");
		Task task = taskRepository.save(beforTask);
		Assert.assertEquals(beforTask, task);
	}
	
	@Test
	public void createTaskTest04() {
		User userBefore = new User("hungnk", "admin123", 0);
		userRepository.save(userBefore);
		Task beforTask = new Task("01", "doing", "hungnk");
		taskService.save(beforTask);
		List<Task> task = taskRepository.findByUserId("hungnk");
		assertFalse(task.isEmpty());
		assertNotNull(task);
	}
	
	@Test
	public void createTask05() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\": \"2022-06-24\" }";
		User userBefore = new User("hungnk", "admin123", 1);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(taskController, "userService", mock);
	    when(mock.loadUserByUsername("hungnk")).thenReturn(userBefore);
	    TaskService taskMock  = Mockito.mock(TaskService.class);
        ReflectionTestUtils.setField(taskController, "taskService", taskMock);
        Task task = new Task("M01", "doning01", "hungnk", LocalDate.now());
        when(taskMock.save(task)).thenReturn(task);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isOk());
	}
	@Test
	public void createTask06() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\": \"2022-06-24\" }";
		User userBefore = new User("hungnk", "admin123", 1);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(taskController, "userService", mock);
	    when(mock.loadUserByUsername("hungnk")).thenReturn(null);
	    TaskService taskMock  = Mockito.mock(TaskService.class);
        ReflectionTestUtils.setField(taskController, "taskService", taskMock);
        Task task = new Task("M01", "doning01", "hungnk", LocalDate.now());
        when(taskMock.save(task)).thenReturn(task);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createTask07() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\": \"2022-06-24\" }";
		User userBefore = new User("hungnk", "admin123", 1);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(taskController, "userService", mock);
	    when(mock.loadUserByUsername("hungnk")).thenReturn(userBefore);
	    TaskService taskMock = Mockito.mock(TaskService.class);
        ReflectionTestUtils.setField(taskController, "taskService", taskMock);
        Task task = new Task("M01", "doning01", "hungnk", LocalDate.now());
        when(taskMock.save(task)).thenReturn(null);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createTask08() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\":\"2022-06-24\" }";
		User userBefore = new User("hungnk", "admin123", 1);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(taskController, "userService", mock);
	    when(mock.loadUserByUsername("hungnk")).thenReturn(userBefore);
	    TaskService taskMock = Mockito.mock(TaskService.class);
        ReflectionTestUtils.setField(taskController, "taskService", taskMock);
        Task task = new Task("M01", "doning01", "hungnk", LocalDate.now());
        Task task2 = new Task("M02", "doning01", "hungnk", LocalDate.now());
        List<Task> tasks = new ArrayList<Task>();
        tasks.add(task);
        tasks.add(task2);
        when(taskMock.findByUserId("hungnk")).thenReturn(tasks);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createTask09() throws Exception {
		String requestJson = "{\"id\":\"M01\",\"content\":\"doning01\",\"userId\":\"hungnk\",\"createdDate\": \"2022-06-24\" }";
		User userBefore = new User("hungnk", "admin123", 1);
		String token = tokenManager.generateJwtToken(userBefore);
		Mockito.when(userServiceMock.loadUserByUsername("hungnk")).thenReturn(userBefore);
		HttpHeaders headers = new HttpHeaders();
	    headers.setBearerAuth(token);
	    UserService mock = Mockito.mock(UserService.class);
        ReflectionTestUtils.setField(taskController, "userService", mock);
	    when(mock.loadUserByUsername("hungnk")).thenReturn(null);
	    TaskService taskMock  = Mockito.mock(TaskService.class);
        ReflectionTestUtils.setField(taskController, "taskService", taskMock);
        Task task = new Task("M01", "doning01", "hungnk", LocalDate.now());
        when(taskMock.save(task)).thenReturn(task);
		mockMvc.perform(
				post("/api/createTask").contentType(APPLICATION_JSON_UTF8).headers(headers).content(requestJson))
				.andExpect(status().isBadRequest());
	}
	
	@Test
	public void createLogin01() throws Exception {
		String login = "{\"username\":\"hungnk\",\"password\":\"admin123\"}";
		mockMvc.perform(post("/api/login").contentType(APPLICATION_JSON_UTF8).content(login))
				.andExpect(status().isOk());
		MessageBody body = new MessageBody(HttpStatus.OK, "Success");
		assertFalse(Objects.isNull(body));
	}
	
}
