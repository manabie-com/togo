package com.manabie.togo.controller;

import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.manabie.togo.jwt.JwtTokenProvider;
import com.manabie.togo.payload.LoginRequest;
import com.manabie.togo.payload.LoginResponse;
import com.manabie.togo.payload.RandomStuff;
import com.manabie.togo.model.CustomUserDetails;
import com.manabie.togo.model.Task;
import com.manabie.togo.payload.AddTaskRequest;
import com.manabie.togo.payload.AddTaskResponse;
import com.manabie.togo.payload.ListTaskRequest;
import com.manabie.togo.payload.ListTaskResponse;
import com.manabie.togo.services.TaskService;
import com.manabie.togo.services.UserService;
import java.util.ArrayList;
import java.util.List;
import javax.servlet.http.HttpServletResponse;
import org.springframework.http.HttpStatus;
import org.springframework.web.server.ResponseStatusException;

/**
 * The controller for all API
 * @author mupmup
 */
@RestController
@RequestMapping("/")
public class RestAPIController {

    @Autowired
    AuthenticationManager authenticationManager;

    @Autowired
    private JwtTokenProvider tokenProvider;

    @Autowired
    private TaskService taskService;

    @Autowired
    private UserService userService;

    /**
     * Login user/pass to app
     * @param loginRequest
     * @return JWT to LoginResponse
     */
    @GetMapping("/login")
    public LoginResponse authenticateUser(@Valid LoginRequest loginRequest) {

        // Xác thực thông tin người dùng Request lên
        Authentication authentication = authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        loginRequest.getUser_id(),
                        loginRequest.getPassword()
                )
        );

        // Nếu không xảy ra exception tức là thông tin hợp lệ
        // Set thông tin authentication vào Security Context
        SecurityContextHolder.getContext().setAuthentication(authentication);

        // Trả về jwt cho người dùng.
        String jwt = tokenProvider.generateToken((CustomUserDetails) authentication.getPrincipal());
        return new LoginResponse(jwt);
    }

    /**
     * List all tasks in database
     * @param request: created_date
     * @return 
     */
    @GetMapping("/tasks")
    public ListTaskResponse listTasks(@Valid ListTaskRequest request) {
        String created_date = request.getCreated_date();
        List<Task> list = taskService.listTaskInDay(created_date);

        ListTaskResponse response = new ListTaskResponse();
        response.setData(list);

        return response;
    }

    /**
     * Add a new content to database
     * @param request
     * @param response
     * @return if under limit then add, if over limit then trigger error Too many requests
     */
    @PostMapping("/tasks")
    public AddTaskResponse addTask(@Valid @RequestBody AddTaskRequest request, HttpServletResponse response) {
        List<Task> list = new ArrayList<>();
        if (request != null) {
            String content = request.getContent();
            Task task = taskService.createNewTask(content);
            if (task != null) {
                list.add(task);
            } else {
                throw new ResponseStatusException(HttpStatus.TOO_MANY_REQUESTS, "Reach the daily limit.");
            }
        }

        AddTaskResponse taskResponse = new AddTaskResponse();
        taskResponse.setData(list);
        return taskResponse;

    }

}
