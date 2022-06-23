package com.interview.challenges.service;

import java.util.List;

import com.interview.challenges.domain.Task;

public interface TaskService {
	Task save(Task task);
	List<Task> findByUserId(String userId);
}
