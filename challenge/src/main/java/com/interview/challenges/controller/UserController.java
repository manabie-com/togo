package com.interview.challenges.controller;

import java.util.Objects;

import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.interview.challenges.domain.User;
import com.interview.challenges.service.UserService;
import com.interview.challenges.utils.CommonUtils;
import com.interview.challenges.utils.MessageBody;

@RestController
@RequestMapping("api")
public class UserController {
	
	@Autowired
	private UserService userService;
	
	@PostMapping(value = "createUser", produces = { "application/json"})
	public ResponseEntity<?> createUser(@Valid @RequestBody User user){
		User saveUser = userService.save(user);
		MessageBody messageBody = new MessageBody();
		if(Objects.nonNull(saveUser)) {
			messageBody.setHttpStatus(HttpStatus.OK);
			messageBody.setMessage(CommonUtils.SUCCESS);
			messageBody.putAll();
			return new ResponseEntity<Object>(messageBody, HttpStatus.OK);
		}
		messageBody.setHttpStatus(HttpStatus.BAD_REQUEST);
		messageBody.setMessage(CommonUtils.FAILD);
		messageBody.putAll();
		return new ResponseEntity<Object>(messageBody, HttpStatus.BAD_REQUEST);
	}
}
