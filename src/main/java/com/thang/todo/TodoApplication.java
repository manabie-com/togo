package com.thang.todo;


import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class TodoApplication {//implements CommandLineRunner {

	public static void main(String[] args) {
		
		SpringApplication.run(TodoApplication.class, args);
	}

	// @Autowired
    // UserRepository userRepository;
	
    // @Autowired
    // PasswordEncoder passwordEncoder;


	// @Override
	// public void run(String... args) throws Exception {
	// 	User user = new User();
    //     user.setUsername("loda");
    //     user.setPassword(passwordEncoder.encode("loda"));
	// 	user.setMaximumTasks(5L);
    //     userRepository.save(user);
    //     System.out.println(user);
	// }

}
