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
package com.manabie.todo.services;

import com.manabie.todo.entity.Todo;
import com.manabie.todo.entity.UserTodo;
import com.manabie.todo.exception.DataNotFoundException;
import com.manabie.todo.model.payload.request.TodoRequest;
import com.manabie.todo.model.payload.request.TodoModelUpdateRequest;
import com.manabie.todo.model.payload.request.UserTodoRequest;
import com.manabie.todo.model.payload.response.TodoResponse;
import com.manabie.todo.model.payload.response.UserTodoResponse;
import com.manabie.todo.repository.TodoRepository;
import com.manabie.todo.repository.UserTodoRepository;
import io.reactivex.Observable;
import org.apache.commons.beanutils.BeanUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Calendar;
import java.util.Date;
import java.util.List;

/**
 *
 * @author :longtran
 * @version :1.0
 */
@Service
public class TodoServiceImpl implements TodoService {

    private final Logger LOGGER = LoggerFactory.getLogger(TodoServiceImpl.class);

    private TodoRepository repository;
    private UserTodoRepository userTodoRepository;

    @Autowired
    public TodoServiceImpl(TodoRepository repository, UserTodoRepository userTodoRepository) {
        this.repository = repository;
        this.userTodoRepository = userTodoRepository;
    }

    /**
     * @param payload
     * @return
     */
    private TodoResponse getResponse(Todo payload) {
        TodoResponse modelResponse = new TodoResponse();
        try {
            BeanUtils.copyProperties(modelResponse, payload);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return modelResponse;
    }

    /**
     *
     * @param payload
     * @return
     */
    private UserTodoResponse getUserTodoResponse(UserTodo payload) {
        UserTodoResponse userTodoResponse = new UserTodoResponse();
        try {
            BeanUtils.copyProperties(userTodoResponse, payload);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return userTodoResponse;
    }

    /**
     *
     * @return
     */
    private Date getEndOfDay() {
        Calendar calendar = Calendar.getInstance();
        int year = calendar.get(Calendar.YEAR);
        int month = calendar.get(Calendar.MONTH);
        int day = calendar.get(Calendar.DATE);
        calendar.setTimeInMillis(0);
        calendar.set(year, month, day, 23, 59, 59);
        return calendar.getTime();
    }

    @Override
    public Observable<UserTodoResponse> createDataLimitWarning(UserTodoRequest payload) {
        return Observable.fromCallable(() -> {
            UserTodo userTodo = new UserTodo();
            userTodo.setUserId(payload.getUserId());
            userTodo.setMaximumTaskPerDay(payload.getMaximumTaskPerDay());
            return userTodoRepository.create (userTodo).map(this::getUserTodoResponse)
                    .orElseThrow(() -> new DataNotFoundException("Create Exception"));
        });
    }

    @Override
    @Transactional
    public Observable<TodoResponse> create(TodoRequest payload) {
        return Observable.fromCallable(() -> {
            UserTodo userTodo = userTodoRepository.getBy(payload.getCreateByUserId())
                    .orElseThrow(() -> new DataNotFoundException("User Not Found"));
            Integer totalTaskOfUser = repository.getTaskBy(payload.getCreateByUserId()).size();
            if(System.currentTimeMillis() < getEndOfDay().getTime() && userTodo.getMaximumTaskPerDay() <= totalTaskOfUser){
                throw new DataNotFoundException("Data limit Exceeded Notification");
            }
            Todo todo = new Todo();
            try {
                BeanUtils.copyProperties(todo, payload);
            } catch (Exception exception) {
                exception.printStackTrace();
            }
            return repository.create(todo).map(this::getResponse)
                    .orElseThrow(() -> new DataNotFoundException("Create Exception"));
        });
    }

    @Override
    @Transactional
    public Observable<TodoResponse> update(TodoModelUpdateRequest payload) {
        return Observable.fromCallable(() -> {
            Todo model = new Todo();
            try {
                BeanUtils.copyProperties(model, payload);
            } catch (Exception exception) {
                exception.printStackTrace();
            }
            return repository.update(model).map(this::getResponse)
                    .orElseThrow(() -> new DataNotFoundException("Update Exception"));
        });
    }

    @Override
    public Observable<TodoResponse> getModel(String id) {
        return Observable.fromCallable(() -> {
            return repository.get(id).map(this::getResponse)
                    .orElseThrow(() -> new DataNotFoundException("Data Not Found Exception"));
        });
    }

    @Override
    public Observable<List<Todo>> getAll(Pageable pageable) {
        return Observable.fromCallable(() -> {
            return repository.getAll(pageable);
        });
    }
}
