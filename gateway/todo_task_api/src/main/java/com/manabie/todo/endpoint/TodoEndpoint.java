/*
 * Copyright (c) 2022, 2022 manabie.com and/or its affiliates. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package com.manabie.todo.endpoint;

import com.manabie.todo.model.payload.request.TodoRequest;
import com.manabie.todo.model.payload.request.UserTodoRequest;
import com.manabie.todo.model.payload.response.TodoResponse;
import com.manabie.todo.services.TodoService;
import io.reactivex.Observable;
import io.reactivex.schedulers.Schedulers;
import io.swagger.annotations.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.BindingResult;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.context.request.async.DeferredResult;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.Locale;
import java.util.UUID;

/**
 * @author longtran
 * @version 1.0
 * @created_by longtran
 */
@RestController
@RequestMapping(value = "/api/v1/to-do")
@Api(value = "TodoEndpoint")
public class TodoEndpoint extends AbstractEndpoint {


    private TodoService modelService;

    @Autowired
    public TodoEndpoint(TodoService modelService) {
        this.modelService = modelService;
    }


    @PostMapping(value = "/data-limit-warning")
    @ApiOperation(
            value = "Data limit Exceeded Notification.",
            notes = "Data limit Exceeded Notification.",
            response = TodoResponse.class)
    @ApiResponses(value = {
            @ApiResponse(code = 201, message = "Created", reference = "Role successfully created.", response = TodoResponse.class),
            @ApiResponse(code = 400, message = "Bad Request", reference = "Requested action could not be understood by the system.", response = ResponseEntity.class),
            @ApiResponse(code = 401, message = "Unauthorized", reference = "Requested action requires authentication.", response = ResponseEntity.class),
            @ApiResponse(code = 403, message = "Forbidden", reference = "System refuses to fulfill the requested action.", response = ResponseEntity.class),
            @ApiResponse(code = 500, message = "Internal Server Error", reference = "A generic error has occurred on the system.", response = ResponseEntity.class)
    })
    public DeferredResult<ResponseEntity> createDataLimitWarning(
            @RequestBody @Validated UserTodoRequest payload,
            BindingResult bindingResult,
            HttpServletRequest httpServletRequest,
            HttpServletResponse httpServletResponse) {
        return toDeferredResult(
                new DeferredResult<>(),
                toObservable(bindingResult).flatMap(
                        to -> modelService.createDataLimitWarning(payload)
                                .subscribeOn(Schedulers.io()),
                        Observable::error,
                        Observable::empty
                ).subscribeOn(Schedulers.io()),
                bindingResult,
                Locale.getDefault(),
                httpServletRequest,
                httpServletResponse
        );
    }

    @PostMapping(value = "/task")
    @ApiOperation(
            value = "Create a new task.",
            notes = "Create a new task.",
            response = TodoResponse.class)
    @ApiResponses(value = {
            @ApiResponse(code = 201, message = "Created", reference = "Role successfully created.", response = TodoResponse.class),
            @ApiResponse(code = 400, message = "Bad Request", reference = "Requested action could not be understood by the system.", response = ResponseEntity.class),
            @ApiResponse(code = 401, message = "Unauthorized", reference = "Requested action requires authentication.", response = ResponseEntity.class),
            @ApiResponse(code = 403, message = "Forbidden", reference = "System refuses to fulfill the requested action.", response = ResponseEntity.class),
            @ApiResponse(code = 500, message = "Internal Server Error", reference = "A generic error has occurred on the system.", response = ResponseEntity.class)
    })
    public DeferredResult<ResponseEntity> create(
            @RequestBody @Validated TodoRequest payload,
            BindingResult bindingResult,
            HttpServletRequest httpServletRequest,
            HttpServletResponse httpServletResponse) {
        return toDeferredResult(
                new DeferredResult<>(),
                toObservable(bindingResult).flatMap(
                        to -> modelService.create(payload)
                                .subscribeOn(Schedulers.io()),
                        Observable::error,
                        Observable::empty
                ).subscribeOn(Schedulers.io()),
                bindingResult,
                Locale.getDefault(),
                httpServletRequest,
                httpServletResponse
        );
    }

    @GetMapping(value = "/task/{taskID}")
    @ApiOperation(
            value = "Retrieve a task.",
            notes = "Retrieve a task.",
            response = TodoResponse.class)
    @ApiResponses(value = {
            @ApiResponse(code = 200, message = "OK", reference = "Role successfully retrieved.", response = TodoResponse.class),
            @ApiResponse(code = 400, message = "Bad Request", reference = "Requested action could not be understood by the system.", response = ResponseEntity.class),
            @ApiResponse(code = 401, message = "Unauthorized", reference = "Requested action requires authentication.", response = ResponseEntity.class),
            @ApiResponse(code = 403, message = "Forbidden", reference = "System refuses to fulfill the requested action.", response = ResponseEntity.class),
            @ApiResponse(code = 500, message = "Internal Server Error", reference = "A generic error has occurred on the system.", response = ResponseEntity.class)
    })
    public DeferredResult<ResponseEntity> get(
            @ApiParam(value = "taskID", example = "9781337a-a4f6-4ee9-b7b2-2c001d8d457d") @PathVariable("taskID") UUID taskID,
            HttpServletRequest httpServletRequest,
            HttpServletResponse httpServletResponse) {
        return toDeferredResult(
                new DeferredResult<>(),
                modelService.getModel(taskID.toString()).subscribeOn(Schedulers.io()),
                Locale.getDefault(),
                httpServletRequest,
                httpServletResponse
        );
    }

    @GetMapping(value = "/tasks/")
    @ApiOperation(
            value = "Retrieve list task.",
            notes = "Retrieve list task.",
            response = TodoResponse.class)
    @ApiResponses(value = {
            @ApiResponse(code = 200, message = "OK", reference = "Role successfully retrieved.", response = TodoResponse.class),
            @ApiResponse(code = 400, message = "Bad Request", reference = "Requested action could not be understood by the system.", response = ResponseEntity.class),
            @ApiResponse(code = 401, message = "Unauthorized", reference = "Requested action requires authentication.", response = ResponseEntity.class),
            @ApiResponse(code = 403, message = "Forbidden", reference = "System refuses to fulfill the requested action.", response = ResponseEntity.class),
            @ApiResponse(code = 500, message = "Internal Server Error", reference = "A generic error has occurred on the system.", response = ResponseEntity.class)
    })
    public DeferredResult<ResponseEntity> get(
            @RequestParam(defaultValue = "1", required = false) Integer page,
            @RequestParam(defaultValue = "5", required = false) Integer size,
            HttpServletRequest httpServletRequest,
            HttpServletResponse httpServletResponse) {
        return toDeferredResult(
                new DeferredResult<>(),
                modelService.getAll(PageRequest.of(Math.max(0, page - 1), size)).subscribeOn(Schedulers.io()),
                Locale.getDefault(),
                httpServletRequest,
                httpServletResponse
        );
    }
}
