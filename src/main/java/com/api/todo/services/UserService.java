package com.api.todo.services;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.api.todo.entities.User;
import com.api.todo.repositories.UserRepository;

@Service
public class UserService {
    @Autowired
    private UserRepository userRepository;

    public UserService() {}
    public User createUser(User user) {
        return userRepository.save(user);
    }

    public Optional<User> findById(long id) {
        return userRepository.findById(id);
    }


    public int getLimitTaskOfUser(long userId) {
        Optional<User> userRepo = userRepository.findById(userId);
        return userRepo.isPresent() ? userRepo.get().getLimitTasksPerDay() : 0;
    }
}
