package com.manabie.todotaskapplication.controller.task.impl;

import com.manabie.todotaskapplication.common.constant.Header;
import com.manabie.todotaskapplication.data.model.Task;
import com.manabie.todotaskapplication.data.pojo.task.ResponseListTaskDto;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import com.manabie.todotaskapplication.repository.task.TaskRepository;
import com.manabie.todotaskapplication.repository.userconfig.UserConfigRepository;
import com.manabie.todotaskapplication.service.userconfig.UserConfigService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.http.*;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Arrays;
import java.util.Objects;
import java.util.Optional;
import java.util.UUID;

import static junit.framework.TestCase.assertTrue;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class TaskApiControllerTest_IntegrationTest {

    @LocalServerPort
    private int port;

    @Value("${server.servlet.context-path}")
    private String contextPath;

    TestRestTemplate restTemplate = new TestRestTemplate();

    HttpHeaders headers = new HttpHeaders();

    private String createURLWithPort(String uri) {
        return "http://localhost:" + port + contextPath + uri;
    }

    @MockBean
    private TaskRepository taskRepository;

    @MockBean
    private UserConfigRepository userConfigRepository;

    @MockBean
    private UserConfigService userConfigService;

    @BeforeEach
    public void setUp() {
    }

    @Test
    public void createTask_Success() {
        Optional<Integer> rateLimit = Optional.of(5);
        String userId = "123123123";
        TaskDto taskDto = new TaskDto();
        taskDto.setName("test-1");
        taskDto.setContent("content-test-1");

        when(userConfigService.getRateLimitConfig(userId)).thenReturn(rateLimit);
        when(taskRepository.save(any())).thenReturn(new Task(UUID.randomUUID(), taskDto.getName(), taskDto.getContent(), userId));
        headers.add(Header.AUTHORIZATION.getValue(), "123123123");

        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks"), HttpMethod.POST, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.OK.equals(respHttpStatus));
    }

    @Test
    public void createTask_Failed_Unauthorized() {
        Optional<Integer> rateLimit = Optional.of(5);
        String wrongUserId = "userId";
        TaskDto taskDto = new TaskDto();
        taskDto.setName("test-1");
        taskDto.setContent("content-test-1");

        when(userConfigService.getRateLimitConfig(wrongUserId)).thenReturn(rateLimit);
        when(taskRepository.save(any())).thenReturn(new Task(UUID.randomUUID(), taskDto.getName(), taskDto.getContent(), wrongUserId));
        headers.add(Header.AUTHORIZATION.getValue(), wrongUserId);

        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks"), HttpMethod.POST, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.UNAUTHORIZED.equals(respHttpStatus));
    }

    @Test
    public void createTask_Failed_Rate_Limit_When_User_Config_Not_Set() {
        String rightUserId = "123123123";
        TaskDto taskDto = new TaskDto();
        taskDto.setName("test-1");
        taskDto.setContent("content-test-1");

        when(taskRepository.save(any())).thenReturn(new Task(UUID.randomUUID(), taskDto.getName(), taskDto.getContent(), rightUserId));
        headers.add(Header.AUTHORIZATION.getValue(), rightUserId);

        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks"), HttpMethod.POST, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.TOO_MANY_REQUESTS.equals(respHttpStatus));
    }

    @Test
    public void createTask_Failed_Rate_Limit_When_Exceed_Created_Time() {
        Optional<Integer> rateLimit = Optional.of(5);
        String rightUserId = "123123123";

        TaskDto taskDto = new TaskDto();
        taskDto.setName("test-1");
        taskDto.setContent("content-test-1");

        when(userConfigService.getRateLimitConfig(rightUserId)).thenReturn(rateLimit);
        when(taskRepository.save(any())).thenReturn(new Task(UUID.randomUUID(), taskDto.getName(), taskDto.getContent(), rightUserId));

        headers.add(Header.AUTHORIZATION.getValue(), rightUserId);
        int counter = 1;
        ResponseEntity<String> response = new ResponseEntity<>(HttpStatus.OK);
        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        while (counter <= rateLimit.get() + 1) {
            response = restTemplate.exchange(
                    createURLWithPort("/tasks"), HttpMethod.POST, entity, String.class);
            counter++;
        }

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.TOO_MANY_REQUESTS.equals(respHttpStatus));

    }

    @Test
    public void createTask_Failed_Rate_Limit_When_Failed_Validate_Business_Rules_Empty_Task_Name() {
        Optional<Integer> rateLimit = Optional.of(5);
        String rightUserId = "122322";
        TaskDto invalidTaskDto = getInvalidTaskDto();

        when(userConfigService.getRateLimitConfig(rightUserId)).thenReturn(rateLimit);
        when(taskRepository.save(any())).thenReturn(new Task(UUID.randomUUID(), invalidTaskDto.getName(), invalidTaskDto.getContent(), rightUserId));
        headers.add(Header.AUTHORIZATION.getValue(), rightUserId);

        HttpEntity<TaskDto> entity = new HttpEntity<>(invalidTaskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks"), HttpMethod.POST, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.BAD_REQUEST.equals(respHttpStatus));
    }


    @Test
    public void getTaskById_Success() throws Exception {
        TaskDto taskDto = getValidTaskDto();
        UUID taskId = UUID.randomUUID();

        when(taskRepository.findById(taskId)).thenReturn(Optional.of(new Task(taskId, taskDto.getName(), taskDto.getContent(), "")));

        HttpEntity<TaskDto> entity = new HttpEntity<>(null, headers);
        ResponseEntity<TaskDto> response = restTemplate.exchange(
                createURLWithPort("/tasks/" + taskId), HttpMethod.GET, entity, TaskDto.class);
        assertTrue(HttpStatus.OK.equals(response.getStatusCode()));
        assertTrue(taskId.equals(Objects.requireNonNull(response.getBody()).getId()));
    }

    @Test
    public void getTaskById_Failed_Not_Found() throws Exception {
        UUID taskId = UUID.randomUUID();

        HttpEntity<TaskDto> entity = new HttpEntity<>(null, headers);
        ResponseEntity<TaskDto> response = restTemplate.exchange(
                createURLWithPort("/tasks/" + taskId), HttpMethod.GET, entity, TaskDto.class);
        assertTrue(HttpStatus.NOT_FOUND.equals(response.getStatusCode()));
    }


    @Test
    public void updateTask_Success() {

        TaskDto taskDto = getValidTaskDto();
        UUID taskId = UUID.randomUUID();
        when(taskRepository.updateTask(taskDto.getName(), taskDto.getContent(), taskId)).thenReturn(1);

        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks/" + taskId), HttpMethod.PUT, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.NO_CONTENT.equals(respHttpStatus));
    }

    @Test
    public void updateTask_Faild_Not_Found_Task_Id() {

        TaskDto taskDto = getValidTaskDto();

        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks/" + UUID.randomUUID()), HttpMethod.PUT, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.NOT_FOUND.equals(respHttpStatus));
    }

    @Test
    public void updateTask_Faild_Validate_Business_Rules() {

        TaskDto taskDto = getInvalidTaskDto();

        HttpEntity<TaskDto> entity = new HttpEntity<>(taskDto, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks/" + UUID.randomUUID()), HttpMethod.PUT, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.BAD_REQUEST.equals(respHttpStatus));
    }

    @Test
    public void deleteTask() {
        TaskDto taskDto = getInvalidTaskDto();

        HttpEntity<TaskDto> entity = new HttpEntity<>(null, headers);
        ResponseEntity<String> response = restTemplate.exchange(
                createURLWithPort("/tasks/" + UUID.randomUUID()), HttpMethod.DELETE, entity, String.class);

        HttpStatus respHttpStatus = response.getStatusCode();
        assertTrue(HttpStatus.NO_CONTENT.equals(respHttpStatus));
    }

    @Test
    public void getTasks() {
        TaskDto taskDto = getValidTaskDto();

        when(taskRepository.findAll()).thenReturn(Arrays.asList(new Task(UUID.randomUUID(), taskDto.getName(), taskDto.getContent(), "123232"),
                new Task(UUID.randomUUID(), taskDto.getName(), taskDto.getContent(), "12323223")));

        HttpEntity<String> entity = new HttpEntity<>(null, headers);
        Class<TaskDto> list = null;
        ResponseEntity<ResponseListTaskDto> response = restTemplate.exchange(
                createURLWithPort("tasks"), HttpMethod.GET, entity, ResponseListTaskDto.class);
        assertTrue(HttpStatus.OK.equals(response.getStatusCode()));
        assertTrue(Objects.requireNonNull(response.getBody()).getData().size() == 2);
    }

    @Test
    public void getTasks_Failed_Not_Found() {
        TaskDto taskDto = getValidTaskDto();

        HttpEntity<String> entity = new HttpEntity<>(null, headers);
        Class<TaskDto> list = null;
        ResponseEntity<TaskDto[]> response = restTemplate.exchange(
                createURLWithPort("tasks"), HttpMethod.GET, entity, TaskDto[].class);
        assertTrue(HttpStatus.NOT_FOUND.equals(response.getStatusCode()));
    }

    private TaskDto getValidTaskDto() {
        TaskDto taskDto = new TaskDto();
        taskDto.setName("test-1");
        taskDto.setContent("content-test-1");
        return taskDto;
    }

    private TaskDto getInvalidTaskDto() {
        TaskDto taskDto = new TaskDto();
        taskDto.setContent("content-test-1");
        return taskDto;
    }
}