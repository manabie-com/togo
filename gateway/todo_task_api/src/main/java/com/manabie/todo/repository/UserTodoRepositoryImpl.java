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

import com.manabie.todo.entity.Todo;
import com.manabie.todo.entity.UserTodo;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

/**
 * @author longtran
 * @version 1.0
 */
@Repository
public class UserTodoRepositoryImpl implements UserTodoRepositoryCustom {

    @PersistenceContext
    private EntityManager entityManager;

    @Override
    @Transactional
    public Optional<UserTodo> create(UserTodo model) {
        UserTodo userTodo = getBy(model.getUserId()).orElse(null);
        if (null == userTodo) {
            entityManager.persist(model);
            entityManager.flush();
            entityManager.clear();
        } else {
            entityManager.merge(userTodo);
        }
        return getBy(model.getUserId());
    }

    @Override
    @Transactional
    public Optional<UserTodo> update(UserTodo modelUpdate) {
        return Optional.ofNullable(entityManager.merge(modelUpdate));
    }

    @Override
    public Optional<UserTodo> get(String id) {
        var query = entityManager.createNativeQuery(
                "SELECT r.* FROM user_todo r WHERE r.id=:id", UserTodo.class);
        query.setParameter("id", id);
        var resultList = query.getResultList();
        if (resultList.size() > 0) {
            return Optional.ofNullable((UserTodo) resultList.get(0));
        } else {
            return Optional.empty();
        }
    }

    @Override
    public List<UserTodo> getAll() {
        var query = entityManager.createNativeQuery(
                "SELECT r.* FROM user_todo r", UserTodo.class);
        var resultList = query.getResultList();
        if (resultList.size() > 0) {
            return resultList;
        } else {
            return new ArrayList<>();
        }
    }

    @Override
    public Optional<UserTodo> getBy(String userID) {
        var query = entityManager.createNativeQuery(
                "SELECT r.* FROM user_todo r WHERE r.user_id=:user_id", UserTodo.class);
        query.setParameter("user_id", userID);
        var resultList = query.getResultList();
        if (resultList.size() > 0) {
            return Optional.ofNullable((UserTodo) resultList.get(0));
        } else {
            return Optional.empty();
        }
    }

    @Override
    public Integer remove(String id) {
        var query = entityManager.createNativeQuery("DELETE FROM user_todo r WHERE r.id=:id");
        query.setParameter("id", id);
        return query.executeUpdate();
    }
}
