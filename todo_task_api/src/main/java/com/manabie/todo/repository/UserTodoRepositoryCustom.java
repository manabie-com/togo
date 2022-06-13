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
package com.manabie.todo.repository;

import com.manabie.todo.entity.UserTodo;

import java.util.List;
import java.util.Optional;

/**
 *
 * @author longtran
 * @version 1.0
 */
public interface UserTodoRepositoryCustom {

    /**
     * Creates a new model.
     */
    Optional<UserTodo> create(UserTodo model);

    /**
     * Updates a model with new values.
     *
     * @return boolean true|false
     */
    Optional<UserTodo> update(UserTodo modelUpdate);

    /**
     * Gets the information.
     *
     * @return role details {@link UserTodo}
     */
    Optional<UserTodo> get(String id);

    /**
     * Gets the information.
     *
     * @return role details {@link UserTodo}
     */
    Optional<UserTodo> getBy(String userID);

    /**
     * Gets the information.
     *
     * @return role details {@link UserTodo}
     */
    List<UserTodo> getAll();

    /**
     * Delete an existing model indicated by the role id.
     */
    Integer remove(String id);
}
