package com.example.demo.controller;


import com.example.demo.model.User;
import com.example.demo.repository.TaskRepository;
import com.example.demo.service.UserService;
import com.example.demo.util.JwtTokenUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.websocket.server.PathParam;

@RestController
@RequestMapping("api/tasks")
public class TaskController {

    @Autowired
    private TaskRepository tasksRepository;

    @Autowired
    private JwtTokenUtil jwtTokenUtil;

    @Autowired
    private UserService userService;

    @GetMapping
    public ResponseEntity getTasks(HttpServletRequest request) {
        final String requestTokenHeader = request.getHeader("Authorization");
        String jwtToken = requestTokenHeader.substring(7);

        User user = userService.loadUserByUsername(jwtTokenUtil.getUsernameFromToken(jwtToken));

        return ResponseEntity.ok(tasksRepository.findByUser(user));
    }

    @GetMapping("/{username}")
    public ResponseEntity getTasks(@PathVariable("username") String username) {
        User user = userService.loadUserByUsername(username);

        return ResponseEntity.ok(tasksRepository.findByUser(user));
    }

}
