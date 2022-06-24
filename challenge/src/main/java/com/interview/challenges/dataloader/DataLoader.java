package com.interview.challenges.dataloader;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;

import com.interview.challenges.domain.User;
import com.interview.challenges.repository.UserRepository;

@Component
public class DataLoader implements ApplicationRunner{

	@Autowired
	UserRepository userRep;
	
	@Autowired
	PasswordEncoder passwordEncoder;
	
	@Override
	public void run(ApplicationArguments args) throws Exception {
		userRep.save(new User("hungnk", passwordEncoder.encode("admin123"), 1));
	}
	
}
