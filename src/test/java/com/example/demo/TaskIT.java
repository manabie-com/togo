package com.example.demo;

import com.example.demo.config.JwtAuthenticationEntryPoint;
import com.example.demo.config.JwtRequestFilter;
import com.example.demo.controller.TaskController;
import com.example.demo.model.Task;
import com.example.demo.model.User;
import com.example.demo.model.UserSettings;
import com.example.demo.repository.AuthorityRepository;
import com.example.demo.repository.TaskRepository;
import com.example.demo.repository.UserRepository;
import com.example.demo.repository.UserSettingsRepository;
import com.example.demo.service.UserService;
import com.example.demo.util.JwtTokenUtil;
import org.junit.Assert;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.context.annotation.Bean;
import org.springframework.http.MediaType;
import org.springframework.mock.web.MockServletContext;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import org.springframework.test.util.ReflectionTestUtils;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.servlet.ServletContext;
import javax.sql.DataSource;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.anyString;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;


@ExtendWith(SpringExtension.class)
@AutoConfigureMockMvc
@WebMvcTest(controllers = TaskController.class)
public class TaskIT {
    @TestConfiguration
    static class TaskIntegrationTestingConfiguration {
        @Bean
        public JwtTokenUtil jwtTokenUtil() {
            return new JwtTokenUtil();
        }
    }
    private MockMvc mockMvc;

    @Autowired
    private WebApplicationContext webApplicationContext;

    @MockBean
    private AuthorityRepository authorityRepository;
    @MockBean
    private UserRepository userRepository;
    @MockBean
    private TaskRepository taskRepository;
    @MockBean
    private UserSettingsRepository userSettingsRepository;
    @MockBean
    private DataSource dataSource;
    @MockBean
    private UserService userService;
    @MockBean
    private EntityManager entityManager;
    @MockBean
    private JwtRequestFilter jwtRequestFilter;
    @MockBean
    private JwtAuthenticationEntryPoint jwtAuthenticationEntryPoint;



    @MockBean
    private EntityManagerFactory entityManagerFactory;

    private String jwt;


    @BeforeEach
    public void setup() throws Exception {
        JwtTokenUtil jwtTokenUtil = new JwtTokenUtil();
        Mockito.when(entityManagerFactory.createEntityManager()).thenReturn(entityManager);
        ReflectionTestUtils.setField(jwtRequestFilter, "jwtTokenUtil", jwtTokenUtil);
        this.mockMvc = MockMvcBuilders.webAppContextSetup(this.webApplicationContext).build();
        Map<String, Object> claims = new HashMap<>();
        User user = new User("test", "password");
        jwt = jwtTokenUtil.generateToken(user);
    }

    @Test
    public void test() throws Exception {
        ServletContext servletContext = webApplicationContext.getServletContext();

        Assert.assertNotNull(servletContext);
        Assert.assertTrue(servletContext instanceof MockServletContext);
        Assert.assertNotNull(webApplicationContext.getBean("taskController"));
    }

    @Test
    public void test_200() throws Exception {
        this.mockMvc.perform(get("/api/tasks").header("Authorization", "Bearer "+jwt))
                .andDo(print()).andExpect(status().isOk());
    }

    @Test
    public void test_get_task_200() throws Exception {
        User user = new User("test", "password");
        Task task = new Task();
        task.setTaskDetails("Test");
        task.setUser(user);
        List<Task> tasks = new ArrayList();
        tasks.add(task);
        Mockito.when(userService.loadUserByUsername(anyString())).thenReturn(user);
        Mockito.when(taskRepository.findByUser(any(User.class))).thenReturn(tasks);
        this.mockMvc.perform(get("/api/tasks").header("Authorization", "Bearer "+jwt))
                .andDo(print()).andExpect(content().json("[{\"id\":null,\"taskDetails\":\"Test\"," +
                        "\"user\":{\"username\":\"test\"},\"isCompleted\":null}]"));
    }
    @Test
    public void test_create_task_200() throws Exception {
        User user = new User("test", "password");
        Task task = new Task();
        task.setTaskDetails("Test");
        task.setUser(user);
        task.setIsCompleted(Boolean.FALSE);
        List<Task> tasks = new ArrayList();
        tasks.add(task);
        UserSettings userSettings = new UserSettings();
        userSettings.setDailyLimit(3L);
        userSettings.setUser(user);

        Mockito.when(userService.loadUserByUsername(anyString())).thenReturn(user);
        Mockito.when(taskRepository.findByUser(any(User.class))).thenReturn(tasks);
        Mockito.when(taskRepository.findById(anyLong())).thenReturn(Optional.of(task));
        Mockito.when(userService.getUserSettings(any(User.class))).thenReturn(userSettings);
        this.mockMvc.perform(post("/api/tasks")
                        .header("Authorization", "Bearer " + jwt)
                        .accept(MediaType.APPLICATION_JSON)
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .content("{\n" +
                                "    \"taskDetail\": \"Test 2\",\n" +
                                "    \"isCompleted\": false\n" +
                                "}"))
                .andDo(print()).andExpect(content().json("{\"id\":null,\"taskDetails\":\"Test 2\"," +
                        "\"user\":{\"username\":\"test\"},\"isCompleted\":false}"))
                .andExpect(status().isOk());
    }

    @Test
    public void test_complete_task_200() throws Exception {
        User user = new User("test", "password");
        Task task = new Task();
        task.setTaskDetails("Test");
        task.setUser(user);
        task.setIsCompleted(Boolean.FALSE);
        List<Task> tasks = new ArrayList();
        tasks.add(task);
        Mockito.when(userService.loadUserByUsername(anyString())).thenReturn(user);
        Mockito.when(taskRepository.findByUser(any(User.class))).thenReturn(tasks);
        Mockito.when(taskRepository.findById(anyLong())).thenReturn(Optional.of(task));
        this.mockMvc.perform(put("/api/tasks")
                        .header("Authorization", "Bearer " + jwt)
                        .accept(MediaType.APPLICATION_JSON)
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .content("{\n" +
                                "    \"id\": 1,\n" +
                                "    \"isTaskCompleted\": true\n" +
                                "}"))
                .andDo(print()).andExpect(content().json("{\"id\":null,\"taskDetails\":\"Test\"," +
                        "\"user\":{\"username\":\"test\"},\"isCompleted\":true}"))
                .andExpect(status().isOk());
    }

    @Test
    public void test_complete_task_by_other_user_400() throws Exception {
        User user = new User("test", "password");
        Task task = new Task();
        task.setTaskDetails("Test");
        task.setUser(user);
        task.setIsCompleted(Boolean.FALSE);
        List<Task> tasks = new ArrayList();
        tasks.add(task);
        Mockito.when(userService.loadUserByUsername(anyString())).thenReturn(user);
        Mockito.when(taskRepository.findByUser(any(User.class))).thenReturn(tasks);
        Mockito.when(taskRepository.findById(anyLong())).thenReturn(Optional.of(task));

        User otherUser = new User("other-user", "password");
        jwt = new JwtTokenUtil().generateToken(otherUser);
        this.mockMvc.perform(put("/api/tasks")
                        .header("Authorization", "Bearer " + jwt)
                        .accept(MediaType.APPLICATION_JSON)
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .content("{\n" +
                                "    \"id\": 1,\n" +
                                "    \"isTaskCompleted\": true\n" +
                                "}"))
                .andDo(print()).andExpect(content().json("{\"message\":\"This task does not belong to this user!\"}"))
                .andExpect(status().isBadRequest());
    }

    @Test
    public void test_delete_task_200() throws Exception {
        User user = new User("test", "password");
        Task task = new Task();
        task.setTaskDetails("Test");
        task.setUser(user);
        task.setIsCompleted(Boolean.FALSE);
        List<Task> tasks = new ArrayList();
        tasks.add(task);
        Mockito.when(userService.loadUserByUsername(anyString())).thenReturn(user);
        Mockito.when(taskRepository.findByUser(any(User.class))).thenReturn(tasks);
        Mockito.when(taskRepository.findById(anyLong())).thenReturn(Optional.of(task));

        this.mockMvc.perform(delete("/api/tasks")
                        .header("Authorization", "Bearer " + jwt)
                        .accept(MediaType.APPLICATION_JSON)
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .content("{\n" +
                                "    \"id\": 1\n" +
                                "}"))
                .andDo(print())
                .andExpect(status().isOk());
    }

    @Test
    public void test_delete_task_by_other_user_400() throws Exception {
        User user = new User("test", "password");
        Task task = new Task();
        task.setTaskDetails("Test");
        task.setUser(user);
        task.setIsCompleted(Boolean.FALSE);
        List<Task> tasks = new ArrayList();
        tasks.add(task);
        Mockito.when(userService.loadUserByUsername(anyString())).thenReturn(user);
        Mockito.when(taskRepository.findByUser(any(User.class))).thenReturn(tasks);
        Mockito.when(taskRepository.findById(anyLong())).thenReturn(Optional.of(task));

        User otherUser = new User("other-user", "password");
        jwt = new JwtTokenUtil().generateToken(otherUser);
        this.mockMvc.perform(delete("/api/tasks")
                        .header("Authorization", "Bearer " + jwt)
                        .accept(MediaType.APPLICATION_JSON)
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .content("{\n" +
                                "    \"id\": 1\n" +
                                "}"))
                .andDo(print())
                .andExpect(content().json("{\"message\":\"This task does not belong to this user!\"}"))
                .andExpect(status().isBadRequest());
    }
}
