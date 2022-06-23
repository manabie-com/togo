package com.interview.challenges.service.impl;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.interview.challenges.domain.Task;
import com.interview.challenges.repository.TaskRepository;
import com.interview.challenges.service.TaskService;

@Service
public class TaskServiceImp implements TaskService{
	
	@Autowired
	TaskRepository taskRepository;
	
	@Override
	public Task save(Task task) {
		return taskRepository.save(task);
	}
	
	@Override
	public List<Task> findByUserId(String userId) {
		// TODO Auto-generated method stub
		return taskRepository.findByUserId(userId);
	}
	
	
}
