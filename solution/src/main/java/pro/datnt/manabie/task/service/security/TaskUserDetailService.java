package pro.datnt.manabie.task.service.security;

import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import pro.datnt.manabie.task.model.UserDBO;
import pro.datnt.manabie.task.repository.UserRepository;

import java.util.Collections;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class TaskUserDetailService implements UserDetailsService {
    private final UserRepository userRepository;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        Optional<UserDBO> user = userRepository.findByUsername(username);

        return user
                .map(u -> new User(u.getUsername(), u.getPassword(), Collections.emptyList()))
                .orElseThrow(() -> new UsernameNotFoundException(username + " not found"));
    }
}
