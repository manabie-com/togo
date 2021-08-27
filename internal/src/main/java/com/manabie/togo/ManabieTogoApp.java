package com.manabie.togo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;


/**
 * Main program
 * @author mupmup
 */
@SpringBootApplication
public class ManabieTogoApp  {

    public static void main(String[] args) {
        SpringApplication.run(ManabieTogoApp.class, args);
    }

    /*
    @Autowired
    UserRepository userRepository;

    @Autowired
    PasswordEncoder passwordEncoder;

    @Override
    public void run(String... args) throws Exception {
        User user = new User();
        user.setUsername("firstUser");
        user.setPassword(passwordEncoder.encode("example"));
        user.setMax_todo(5l);
        userRepository.save(user);
        System.out.println(user);
    }*/
 

}
