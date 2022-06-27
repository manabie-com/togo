package com.api.todo.validators;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Locale;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.MessageSource;
import org.springframework.stereotype.Component;
import org.springframework.validation.Errors;
import org.springframework.validation.ValidationUtils;
import org.springframework.validation.Validator;

import com.api.todo.entities.User;
import com.api.todo.request.RequestTaskEntity;
import com.api.todo.services.TodoService;
import com.api.todo.services.UserService;

@Component
public class TaskValidator implements Validator {
	@Autowired
	private MessageSource messageSource;
	@Autowired
	private UserService userService;
	@Autowired
	private TodoService todoService;

	public TaskValidator() {}
	@Override
	public boolean supports(Class<?> clazz) {
		return RequestTaskEntity.class.equals(clazz);
	}

	@Override
	public void validate(Object target, Errors errors) {
		RequestTaskEntity taskEntity = (RequestTaskEntity) target;

		// check empty
		checkEmptyAndNegative(errors, taskEntity);

		// check max length
		checkMaxLength(errors, taskEntity);

		// check userId existed or not
		checkUserIdExistOrNot(taskEntity.getUserId(), errors);

		// check limit tasks of user
		checkLimitTasksOfUser(taskEntity.getUserId(), errors);
	}

	private void checkEmptyAndNegative(Errors errors, RequestTaskEntity taskEntity) {
		ValidationUtils.rejectIfEmpty(errors, "title",
				messageSource.getMessage("task.title.error", null, Locale.getDefault()));
		ValidationUtils.rejectIfEmpty(errors, "description",
				messageSource.getMessage("task.description.error", null, Locale.getDefault()));

		// check negative number
		if (taskEntity.getUserId() <= 0) {
			errors.rejectValue("userId", messageSource.getMessage("task.userId.error", null, Locale.getDefault()),
					"User ID should be greater than zero");
		}
	}

	private void checkMaxLength(Errors errors, RequestTaskEntity taskEntity) {
		if (taskEntity.getTitle() != null && !taskEntity.getTitle().isEmpty() && taskEntity.getTitle().length() > 50) {
			errors.rejectValue("title", messageSource.getMessage("task.title.maxlength", null, Locale.getDefault()),
					"Title should be smaller than 50 characters");
		}

		if (taskEntity.getDescription() != null && !taskEntity.getDescription().isEmpty()
				&& taskEntity.getDescription().length() > 250) {
			errors.rejectValue("description",
					messageSource.getMessage("task.description.maxlength", null, Locale.getDefault()),
					"Title should be smaller than 250 characters");
		}
	}

	private void checkUserIdExistOrNot(long userId, Errors errors) {
		if (!userService.findById(userId).isPresent()) {
			errors.rejectValue("userId", messageSource.getMessage("task.userId.not_existed", null, Locale.getDefault()),
					"userId field should be existed in table user");
		}
	}

	private void checkLimitTasksOfUser(long userId, Errors errors) {
		Optional<User> user = userService.findById(userId);
		if (!user.isPresent()) {
			return;
		}

		LocalDateTime ldt = LocalDateTime.now();
		System.out.println(DateTimeFormatter.ofPattern("yyyy-MM-dd", Locale.ENGLISH).format(ldt));
		int numTaskPerDay = todoService.countTaskOfOneUser(userId,
				DateTimeFormatter.ofPattern("yyyy-MM-dd", Locale.ENGLISH).format(ldt));
		if (user.get().getLimitTasksPerDay() == 0
				|| (user.get().getLimitTasksPerDay() > 0 && numTaskPerDay >= user.get().getLimitTasksPerDay())) {
			errors.rejectValue("userId", messageSource.getMessage("task.userId.reject_task", null, Locale.getDefault()),
					"userId field enough task");
		}

	}
}