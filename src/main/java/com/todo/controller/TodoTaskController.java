package com.todo.controller;

import com.todo.entity.AppAccount;
import com.todo.entity.TodoTask;
import com.todo.model.TodoTaskDTO;
import com.todo.service.account.AccountService;
import com.todo.service.task.TodoTaskService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/tasks")
public class TodoTaskController {
    protected final Log logger = LogFactory.getLog(getClass());

    @Autowired
    AccountService appAccountService;

    @Autowired
    TodoTaskService todoTaskService;

    @GetMapping
    public ResponseEntity<List<TodoTask>> findAll() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        return ResponseEntity.ok(todoTaskService.findTodoTaskByAppAccount_Username(authentication.getName()));
    }

    @PostMapping()
    public ResponseEntity<TodoTask> save(@RequestBody String content) {
        logger.warn("content : " + content);

        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        String name = authentication.getName();
        logger.warn("authentication : " + name);

        try {
            AppAccount appUserByEmail = appAccountService.findByUsername(name);
            logger.warn("appUserByEmail : " + appUserByEmail);

            boolean limitedTask = isLimitedTask(appUserByEmail.getUsername(), appUserByEmail.getUserTaskLimit());
            logger.warn("limitedTask : " + limitedTask);

            if (limitedTask) {
                return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);

            } else {
                TodoTask todoTask = new TodoTask(content);
                todoTask.setAppAccount(appUserByEmail);
                logger.warn("todoTask : " + todoTask);

                return ResponseEntity.ok(todoTaskService.create(todoTask));
            }
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }

    }

    @RequestMapping(value = "{taskID}", method = RequestMethod.POST, produces = {MediaType.APPLICATION_JSON_VALUE})
    public ResponseEntity<TodoTask> update(@PathVariable("taskID") Long id, @RequestBody TodoTaskDTO todoTaskDTO) {
        return ResponseEntity.ok(todoTaskService.update(id, todoTaskDTO));
    }

    private boolean isLimitedTask(String username, int taskLimit) {
        logger.warn("username : " + username);
        logger.warn("taskLimit : " + taskLimit);
        int i = todoTaskService.countByAppAccount_Username(username);
        logger.warn("i : " + i);

        return i >= taskLimit;
    }

}
