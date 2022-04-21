package com.manabie.interview.task.service;

import com.manabie.interview.task.model.User;
import com.manabie.interview.task.repository.UserRepository;
import com.manabie.interview.task.response.APIResponse;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
@AllArgsConstructor
@Service
public class UserService implements UserDetailsService {
    private final UserRepository userRepository;
    private final BCryptPasswordEncoder bCryptPasswordEncoder;


    public APIResponse registerNewUser(User user) {
        boolean existsById = userRepository.existsById(user.getUid());

        if(existsById){
            return new APIResponse("UserID existed", HttpStatus.OK);
        }else {
            String encode = bCryptPasswordEncoder.encode(user.getPassword());
            user.setPassword(encode);
            userRepository.save(user);
            return new APIResponse("Successfully Registration", HttpStatus.OK);
        }
    }

    public APIResponse deleteUser(String uid) {
        boolean existsById = userRepository.existsById(uid);
        if(!existsById){
            return new APIResponse("UserID not existed", HttpStatus.OK);
        }else {
            userRepository.deleteById(uid);
            return new APIResponse("Successfully Remove", HttpStatus.OK);
        }
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        Optional<User> user = userRepository.findUserById(username);
        if(user.isPresent()){
            return user.get();
        }
        throw new UsernameNotFoundException(String.format("User %s not existed", username));
    }
}
