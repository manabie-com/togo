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
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

/**
 * @author longtran
 * @version 1.0
 */
@Repository
public class TodoRepositoryImpl implements TodoRepositoryCustom {

    @PersistenceContext
    private EntityManager entityManager;

    @Override
    @Transactional
    public Optional<Todo> create(Todo model) {
        model.setDateCreated(System.currentTimeMillis());
        entityManager.persist(model);
        entityManager.flush();
        entityManager.clear();
        return Optional.ofNullable(model);
    }

    @Override
    @Transactional
    public Optional<Todo> update(Todo modelUpdate) {
        modelUpdate.setDateCreated(System.currentTimeMillis());
        return Optional.ofNullable(entityManager.merge(modelUpdate));
    }

    @Override
    public Optional<Todo> get(String id) {
        var query = entityManager.createNativeQuery(
                "SELECT r.* FROM todo r WHERE r.id=:id", Todo.class);
        query.setParameter("id", id);
        var resultList = query.getResultList();
        if (resultList.size() > 0) {
            return Optional.ofNullable((Todo) resultList.get(0));
        } else {
            return Optional.empty();
        }
    }

    @Override
    public List<Todo> getTaskBy(String userID) {
        var query = entityManager.createNativeQuery(
                "SELECT r.* FROM todo r WHERE r.create_by_user_id=:create_by_user_id", Todo.class);
        query.setParameter("create_by_user_id", userID);
        var resultList = query.getResultList();
        if (resultList.size() > 0) {
            return resultList;
        } else {
            return new ArrayList<>();
        }
    }

    @Override
    public List<Todo> getAll(Pageable pageable) {
        var query = entityManager.createNativeQuery(
                "SELECT r.* FROM todo r", Todo.class);
        query.setFirstResult((int) pageable.getOffset());
        query.setMaxResults(pageable.getPageSize());
        var resultList = query.getResultList();
        if (resultList.size() > 0) {
            return resultList;
        } else {
            return new ArrayList<>();
        }
    }

    @Override
    public Integer remove(String id) {
        var query = entityManager.createNativeQuery("DELETE FROM todo r WHERE r.id=:id");
        query.setParameter("id", id);
        return query.executeUpdate();
    }
}
