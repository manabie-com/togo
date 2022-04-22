package com.manabie.interview.task.controller;

import com.manabie.interview.task.model.User;
import com.manabie.interview.task.model.UserRole;
import com.manabie.interview.task.response.APIResponse;
import com.manabie.interview.task.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(path = "api/user")
public class UserController {

    private final UserService userService;

    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }


    @PostMapping(path = "/register")
    public ResponseEntity<APIResponse> registerNewUser(@RequestParam String uid, @RequestParam String password){
        APIResponse response = userService.registerNewUser(
                new User(uid,
                        password,
                        3,
                        UserRole.USER)
        );
        return new ResponseEntity<>(response, response.getStatus());
    }

    @DeleteMapping(path = "/delete/{userId}")
    public ResponseEntity<APIResponse> deleteUser(@PathVariable("userId") String uid){
        APIResponse response = userService.deleteUser(uid);
        return new ResponseEntity<>(response, response.getStatus());
    }
}
