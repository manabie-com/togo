package com.example.demo.config;

import com.example.demo.model.MyAuthorities;
import com.example.demo.model.Task;
import com.example.demo.model.User;
import com.example.demo.model.UserSettings;
import com.example.demo.repository.AuthorityRepository;
import com.example.demo.repository.TaskRepository;
import com.example.demo.repository.UserRepository;
import com.example.demo.repository.UserSettingsRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.crypto.password.PasswordEncoder;

@Configuration
public class InitConfig {

    @Autowired
    private AuthorityRepository authorityRepository;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private TaskRepository taskRepository;

    @Autowired
    private UserSettingsRepository userSettingsRepository;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Autowired
    private UserDetailsService userDetailsService;

    @Bean
    public void initializer() {
        MyAuthorities userAuth = userAuth();
        MyAuthorities adminAuth = adminAuth();
        authorityRepository.save(userAuth);
        authorityRepository.save(adminAuth());
        String password = passwordEncoder.encode("password");
        User user = new User("florante", password);
        user.addAuthority(adminAuth);
        userRepository.save(user);
        Task task = new Task();
        task.setTaskDetails("Task 1");
        task.setUser(user);
        taskRepository.save(task);
        userSettingsRepository.save(new UserSettings(user, Long.valueOf(3)));


    }

    @Bean
    public MyAuthorities adminAuth() {
        return new MyAuthorities("admin", "ROLE_ADMIN");
    }

    @Bean
    public MyAuthorities userAuth() {
        return new MyAuthorities("user", "ROLE_USER");
    }


}
