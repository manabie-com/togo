package com.antulev.togo.controllers;

import java.security.Principal;
import java.time.LocalDate;
import java.time.ZoneId;
import java.util.Date;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.security.access.prepost.PostFilter;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

import com.antulev.togo.models.Task;
import com.antulev.togo.models.UserInfo;
import com.antulev.togo.repositories.TaskRepository;
import com.antulev.togo.repositories.UserInfoRepository;

@RestController
@RequestMapping("/api/tasks")
public class TaskController {

	@Autowired
	TaskRepository taskRepository;

	@Autowired
	UserInfoRepository userInfoRepository;

	@GetMapping("")
	@PostFilter("filterObject.createdBy == authentication.principal")
	public List<Task> findAll() {
		return taskRepository.findAll();
	}

	@PostMapping("")
	public Task createTask(@RequestBody Map<String, ?> taskData, Principal principal) {
		UserInfo userInfor = userInfoRepository.findByUid(principal.getName());

		List<Task> tasks = taskRepository.findByCreatedDateBetween(
				Date.from(LocalDate.now().atStartOfDay().atZone(ZoneId.systemDefault()).toInstant()),
				Date.from(LocalDate.now().plusDays(1).atStartOfDay().atZone(ZoneId.systemDefault()).toInstant()));
		if (userInfor.getTaskLimit() <= tasks.size()) {
			throw new ResponseStatusException(HttpStatus.LOCKED, "Task number reach limit");
		}
		
		String title;
		String notes;
		try {
			title = (String) taskData.get("title");
			notes = (String) taskData.get("notes");
		}catch (Exception e) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		if(title == null || notes == null) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		if(title.isEmpty() || notes.isEmpty()) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		var task = new Task();
		task.setTitle(title);
		task.setNotes(notes);
		
		return taskRepository.save(task);
	}
	
	@PutMapping("/{id}")
	public Task updateTask(@PathVariable long id, @RequestBody Map<String, ?> taskData) {

		String title;
		String notes;
		try {
			title = (String) taskData.get("title");
			notes = (String) taskData.get("notes");
		}catch (Exception e) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		
		var task = taskRepository.getById(id);
		if(!title.isEmpty()) {
			task.setTitle(title);
		}
		if(!notes.isEmpty()) {
			task.setNotes(notes);
		}
		
		return taskRepository.save(task);
	}
	@PutMapping("/{id}/complete")
	public Task complete(@PathVariable long id) {
		var task = taskRepository.getById(id);
		task.setCompletedDate(new Date());
		return task;
	}
	
	@PutMapping("/{id}/delete")
	public Task delete(@PathVariable long id) {
		var task = taskRepository.getById(id);
		task.setDeletedDate(new Date());
		return task;
	}
	
	@GetMapping("/{id}")
	public Optional<Task> getOne(@PathVariable long id) {
		return taskRepository.findById(id);
	}

}
