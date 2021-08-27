package com.manabie.togo.services;

import com.manabie.togo.model.CustomUserDetails;
import com.manabie.togo.model.User;
import com.manabie.togo.repository.UserRepository;
import javax.transaction.Transactional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

/**
 * Service layer for user
 *
 * @author mupmup
 */
@Service
public class UserService implements UserDetailsService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    PasswordEncoder passwordEncoder;

    public void addUser(String username, String password, long max_todo) throws Exception {
        User user = new User();
        user.setUsername(username);
        user.setPassword(passwordEncoder.encode(password));
        user.setMax_todo(max_todo);
        userRepository.save(user);
    }
    
    public void removeUser(String username){
        userRepository.deleteById(username);
    }
    
    @Override
    public UserDetails loadUserByUsername(String username) {
        // check user exists or not
        User user = userRepository.findByUsername(username);
        if (user == null) {
            throw new UsernameNotFoundException(username);
        }
        return new CustomUserDetails(user);
    }

    // JWTAuthenticationFilter will use this function
    @Transactional
    public UserDetails loadUserById(String id) {
        User user = userRepository.findById(id).orElseThrow(
                () -> new UsernameNotFoundException("User not found with id : " + id)
        );

        return new CustomUserDetails(user);
    }

    /**
     * Load user from database
     *
     * @param id
     * @return
     */
    public User loadUser(String id) {
        User user = userRepository.findById(id).orElse(null);
        return user;
    }

}
