package com.antulev.togo.controllers;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import javax.annotation.security.RolesAllowed;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

import com.antulev.togo.models.Account;
import com.antulev.togo.models.UserInfo;
import com.antulev.togo.repositories.AccountRepository;
import com.antulev.togo.repositories.UserInfoRepository;

@RestController
@RequestMapping("/api/accounts")
public class AccountControler {

	@Autowired
	AccountRepository accountRepository;

	@Autowired
	UserInfoRepository userInfoRepository;

	@GetMapping("")
	@RolesAllowed("ADMIN")
	public List<UserInfo> findAll() {
		return userInfoRepository.findAll();
	}

	@RolesAllowed("ADMIN")
	@GetMapping("/{id}")
	public Optional<UserInfo> getById(@PathVariable long id) {
		return userInfoRepository.findById(id);
	}

	@RolesAllowed("ADMIN")
	@PutMapping("/{id}")
	public UserInfo update(@PathVariable long id, @RequestBody Map<String, ?> userMap) {
		String firstName;
		String lastName;
		int taskLimit;
		String password;
		try {
			firstName = (String) userMap.get("firstName");
			lastName = (String) userMap.get("lastName");
			taskLimit = (int) userMap.get("taskLimit");
			password = (String) userMap.get("password");
		} catch (Exception e) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}

		UserInfo userInfo = userInfoRepository.getById(id);
		if (!firstName.isEmpty()) {
			userInfo.setFirstName(firstName);

		}
		if (!lastName.isEmpty()) {
			userInfo.setLastName(lastName);

		}
		if (taskLimit > 0) {
			userInfo.setTaskLimit(taskLimit);

		}

		var check = accountRepository.findByUid(userInfo.getUid());
		if (check.isEmpty()) {
			throw new ResponseStatusException(HttpStatus.CONFLICT, "This username is not existing");
		}
		Account account = check.get();
		if (!firstName.isEmpty()) {
			account.setFirstName(firstName);

		}
		if (!lastName.isEmpty()) {
			account.setLastName(lastName);

		}
		if (!password.isEmpty()) {
			account.setPassword(digestSHA(password));
		}

		try {
			accountRepository.save(account);
		} catch (Exception e) {
			throw new ResponseStatusException(HttpStatus.CONFLICT, "Cannot reuse password");
		}
		return userInfoRepository.save(userInfo);
	}

	@PostMapping("")
	@RolesAllowed("ADMIN")
	public UserInfo createUser(@RequestBody Map<String, ?> userMap) {

		String firstName;
		String lastName;
		int taskLimit;
		String uid;
		String password;
		try {
			firstName = (String) userMap.get("firstName");
			lastName = (String) userMap.get("lastName");
			taskLimit = (int) userMap.get("taskLimit");
			uid = (String) userMap.get("uid");
			password = (String) userMap.get("password");
		} catch (Exception e) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		if (firstName == null || lastName == null || uid == null || password == null || taskLimit < 1) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		if (firstName.isEmpty() || lastName.isEmpty() || uid.isEmpty() || password.isEmpty() || taskLimit < 1) {
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED, "Wrong input");
		}
		var check = accountRepository.findByUid((String) userMap.get("uid"));
		if (check.isPresent()) {
			throw new ResponseStatusException(HttpStatus.CONFLICT, "This username is already existing");
		}

		UserInfo userInfo = new UserInfo();
		userInfo.setFirstName(firstName);
		userInfo.setLastName(lastName);
		userInfo.setTaskLimit(taskLimit);
		userInfo.setUid(uid);

		Account account = new Account();
		account.setFirstName(firstName);
		account.setLastName(lastName);
		account.setUid(uid);
		account.setPassword(digestSHA(password));

		accountRepository.save(account);

		return userInfoRepository.save(userInfo);
	}

	private String digestSHA(final String password) {
		String base64;
		try {
			MessageDigest digest = MessageDigest.getInstance("SHA");
			digest.update(password.getBytes());
			base64 = Base64.getEncoder().encodeToString(digest.digest());
		} catch (NoSuchAlgorithmException e) {
			throw new RuntimeException(e);
		}
		return "{SHA}" + base64;
	}

}
