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
import com.manabie.todo.model.payload.request.TodoRequest;
import com.manabie.todo.model.payload.request.TodoModelUpdateRequest;
import com.manabie.todo.model.payload.request.UserTodoRequest;
import com.manabie.todo.model.payload.response.TodoResponse;
import com.manabie.todo.model.payload.response.UserTodoResponse;
import io.reactivex.Observable;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 *
 * @author :longtran
 * @version :1.0
 */

public interface TodoService {

    /**
     * Creates a new model.
     */
    Observable<TodoResponse> create(TodoRequest payload);

    /**
     * Updates a model with new values.
     *
     * @return boolean true|false
     */
    Observable<TodoResponse> update(TodoModelUpdateRequest payload);

    /**
     * Gets the information for a model.
     *
     * @return role details {@link TodoResponse}
     */
    Observable<TodoResponse> getModel(String id);


    /**
     * Creates a new model.
     */
    Observable<UserTodoResponse> createDataLimitWarning(UserTodoRequest payload);

    /**
     * Gets the information for a model.
     *
     * @return role details {@link Todo}
     */
    Observable<List<Todo>> getAll(Pageable pageable);
}
