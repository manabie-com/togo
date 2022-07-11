package com.manabie.todotask.service;

import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.entity.Task;

public interface TaskService {
    Task addTask(AddTaskRequest request);
}
