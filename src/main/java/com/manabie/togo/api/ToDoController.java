package com.manabie.togo.api;

import com.manabie.togo.dto.ToDoRequest;
import com.manabie.togo.exception.DailyLimitException;
import com.manabie.togo.exception.UserNotFoundException;
import com.manabie.togo.service.ToDoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.validation.Valid;

@RestController
@RequestMapping("/todo")
public class ToDoController {

    @Autowired
    private ToDoService toDoService;

    @PostMapping(path = "/add",produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<String> addTask(@Valid @RequestBody ToDoRequest toDoRequest) {
        try {
            return ResponseEntity.status(HttpStatus.OK).body(toDoService.addToDo(toDoRequest));
        } catch (UserNotFoundException ex){
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(ex.getMessage());
        } catch (DailyLimitException ex){
            return ResponseEntity.status(HttpStatus.TOO_MANY_REQUESTS).body(ex.getMessage());
        }
    }
}
