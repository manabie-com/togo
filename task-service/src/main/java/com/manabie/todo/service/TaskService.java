package com.manabie.todo.service;

import com.manabie.todo.model.CreateTaskRequest;
import com.manabie.todo.model.TaskModel;
import com.manabie.todo.model.UserInfo;

import java.util.List;

public interface TaskService {
    TaskModel createTask(CreateTaskRequest request);

    List<TaskModel> findAllTask();

    List<TaskModel> findAllByOwner(final Long userId);

    TaskModel getTaskById(Long id);

    UserInfo getUserInfo(Long userId);

}
