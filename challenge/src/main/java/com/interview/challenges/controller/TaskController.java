package com.interview.challenges.controller;

import java.time.LocalDate;
import java.util.List;
import java.util.Objects;

import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.interview.challenges.domain.Task;
import com.interview.challenges.domain.User;
import com.interview.challenges.service.TaskService;
import com.interview.challenges.service.UserService;
import com.interview.challenges.utils.CommonUtils;
import com.interview.challenges.utils.MessageBody;

@RestController
@RequestMapping("api")
public class TaskController {

	@Autowired
	private TaskService taskService;

	@Autowired
	private UserService userService;

	@PostMapping(value = "createTask", produces = { "application/json" })
	public ResponseEntity<?> createUser(@Valid @RequestBody Task task) throws Exception {
		MessageBody messageBody = new MessageBody();
		User user = (User) userService.loadUserByUsername(task.getUserId());
		if (Objects.nonNull(user)) {
			List<Task> tasks = taskService.findByUserId(task.getUserId());
			if (isMaxLimitCreateTask(tasks, user.getMaxLimitTodo(), task)) {
				Task saveTask = taskService.save(task);
				if (Objects.nonNull(saveTask)) {
					messageBody.setHttpStatus(HttpStatus.OK);
					messageBody.setMessage(CommonUtils.SUCCESS);
					messageBody.putAll();
					return ResponseEntity.ok(messageBody);
				}
			}else {
				throw new IllegalAccessException(String.format(CommonUtils.OVER_LIMIT, user.getMaxLimitTodo()));
			}
		}
		throw new Exception(CommonUtils.FAILD);
	}

	private boolean isMaxLimitCreateTask(List<Task> tasks, int maxLimit, Task task) {
		int countLimit = 0;
		if(tasks.size() == 1 && tasks.get(0).getId().equals(task.getId())) {
			return true;
		}
		if(!task.getCreatedDate().isEqual(LocalDate.now())) {
			return true;
		}
		for (Task ts : tasks) {
			if(ts.getCreatedDate().isEqual(LocalDate.now())) {
				countLimit++;
			}
		}
		return countLimit < maxLimit;
	}
}
