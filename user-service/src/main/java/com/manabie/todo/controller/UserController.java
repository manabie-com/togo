package com.manabie.todo.controller;

import com.manabie.todo.model.BaseResponse;
import com.manabie.todo.model.CreateUserRequest;
import com.manabie.todo.model.UserInfo;
import com.manabie.todo.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequiredArgsConstructor
@RequestMapping("/user")
public class UserController {
    final UserService userService;

    @PostMapping
    BaseResponse<UserInfo> createUser(@RequestBody CreateUserRequest request) {
        return BaseResponse.ofSucceeded(userService.create(request));
    }

    @GetMapping("/{userId}")
    BaseResponse<UserInfo> getUserTaskById(@PathVariable Long userId) {
        return BaseResponse.ofSucceeded(userService.getById(userId));
    }

    @GetMapping()
    BaseResponse<List<UserInfo>> getAllUser() {
        return BaseResponse.ofSucceeded(userService.findAll());
    }
}