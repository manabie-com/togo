package com.interview.challenges.service;

import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;

import com.interview.challenges.domain.User;

public interface UserService extends UserDetailsService{
	
	User save(User user);
	
	@Override
	UserDetails loadUserByUsername(String id)throws UsernameNotFoundException;
}
