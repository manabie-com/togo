package pro.datnt.manabie.task.service;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import pro.datnt.manabie.task.controller.model.LoginResponseDTO;
import pro.datnt.manabie.task.controller.model.UserDTO;
import pro.datnt.manabie.task.model.UserDBO;
import pro.datnt.manabie.task.repository.UserRepository;
import pro.datnt.manabie.task.service.security.JwtUtil;

import java.util.Optional;

@RequiredArgsConstructor
@Service
public class UserService {
    private final UserRepository userRepository;
    private final JwtUtil jwtUtil;
    private final PasswordEncoder passwordEncoder;

    public Optional<Long> getUserId(String username) {
        Optional<UserDBO> user = userRepository.findByUsername(username);

        return user.map(UserDBO::getId);
    }

    public Optional<String> loginUser(UserDTO user) {
        return userRepository.findByUsername(user.getUsername())
                .filter((u) -> passwordEncoder.matches(user.getPassword(), u.getPassword()))
                .map((u) -> jwtUtil.generateToken(user.getUsername()));
    }

    public void createUser(UserDTO user) {
        UserDBO dbUser = new UserDBO();
        dbUser.setUsername(user.getUsername());
        dbUser.setPassword(passwordEncoder.encode(user.getPassword()));
        userRepository.save(dbUser);
    }
}
