package com.example.demo.model;


import com.example.demo.exception.TaskException;
import com.example.demo.model.CompleteTaskRequest;
import com.example.demo.model.CreateTaskRequest;
import com.example.demo.model.DeleteTaskRequest;
import com.example.demo.model.Task;
import com.example.demo.model.User;
import com.example.demo.model.UserSettings;
import com.example.demo.repository.TaskRepository;
import com.example.demo.service.UserService;
import com.example.demo.util.JwtTokenUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.websocket.server.PathParam;
import java.util.List;
import java.util.Optional;

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

    @PostMapping
    public ResponseEntity createTask(HttpServletRequest request, @RequestBody CreateTaskRequest createTaskRequest) throws Exception {
        final String requestTokenHeader = request.getHeader("Authorization");
        String jwtToken = requestTokenHeader.substring(7);

        User user = userService.loadUserByUsername(jwtTokenUtil.getUsernameFromToken(jwtToken));


        // Check limit
        UserSettings userSettings = userService.getUserSettings(user);
        List<Task> tasks = tasksRepository.findByUser(user);

        if (tasks.size() >= userSettings.getDailyLimit()) {
            throw new TaskException("User has reached daily limit of tasks");
        }

        Task task = new Task();
        task.setTaskDetails(createTaskRequest.getTaskDetail());
        task.setIsCompleted(createTaskRequest.getIsCompleted());
        task.setUser(user);
        tasksRepository.save(task);
        return ResponseEntity.ok(task);
    }

    @PutMapping
    public ResponseEntity completeTask(HttpServletRequest request, @RequestBody CompleteTaskRequest completeTaskRequest) throws Exception {
        final String requestTokenHeader = request.getHeader("Authorization");
        String jwtToken = requestTokenHeader.substring(7);
        User user = userService.loadUserByUsername(jwtTokenUtil.getUsernameFromToken(jwtToken));
        Optional<Task> opTask = tasksRepository.findById(completeTaskRequest.getId());
        Task task = null;
        if (opTask.isPresent()) {
            task = opTask.get();
        } else {
            throw new TaskException("Task does not exist");
        }
        if (!user.getUsername().equals(task.getUser().getUsername())) {
            throw new TaskException("This task does not belong to this user!");
        }
        task.setIsCompleted(completeTaskRequest.getIsTaskCompleted());
        tasksRepository.save(task);
        return ResponseEntity.ok(task);
    }

    //Admin Endpoint
    @GetMapping("/{username}")
    public ResponseEntity getTasks(@PathVariable("username") String username) {
        User user = userService.loadUserByUsername(username);

        return ResponseEntity.ok(tasksRepository.findByUser(user));
    }

    @DeleteMapping
    public ResponseEntity deleteTask(HttpServletRequest request, @RequestBody DeleteTaskRequest deleteTaskRequest) throws Exception {
        final String requestTokenHeader = request.getHeader("Authorization");
        String jwtToken = requestTokenHeader.substring(7);
        User user = userService.loadUserByUsername(jwtTokenUtil.getUsernameFromToken(jwtToken));
        Optional<Task> opTask = tasksRepository.findById(deleteTaskRequest.getId());
        Task task = null;
        if (opTask.isPresent()) {
            task = opTask.get();
        } else {
            throw new TaskException("Task does not exist");
        }
        if (!user.getUsername().equals(task.getUser().getUsername())) {
            throw new TaskException("This task does not belong to this user!");
        }
        tasksRepository.delete(task);
        return ResponseEntity.ok().build();
    }

}
