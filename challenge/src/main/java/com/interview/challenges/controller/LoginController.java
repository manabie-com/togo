package com.interview.challenges.controller;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.DisabledException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.interview.challenges.domain.User;
import com.interview.challenges.security.jwtutils.TokenManager;
import com.interview.challenges.service.UserService;

@RestController
@RequestMapping("api")
public class LoginController {

	@Autowired
	private UserService userService;

	@Autowired
	private AuthenticationManager authenticationManager;

	@Autowired
	private TokenManager tokenManager;

	@PostMapping("/login")
	public ResponseEntity<?> createToken(@RequestBody User user) throws Exception {
		try {
			authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(user.getUsername(), user.getPassword()));
		} catch (DisabledException e) {
			throw new Exception("USER_DISABLED", e);
		} catch (BadCredentialsException e) {
			throw new Exception("INVALID_CREDENTIALS", e);
		}
		final UserDetails userDetails = userService.loadUserByUsername(user.getUsername());
		final String jwtToken = tokenManager.generateJwtToken(userDetails);
		Map<String, Object> map = new HashMap<String, Object>();
		map.put("token", jwtToken);
		return ResponseEntity.ok(map);
	}
}
