package pro.datnt.manabie.task.controller;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import pro.datnt.manabie.task.controller.model.LoginResponseDTO;
import pro.datnt.manabie.task.controller.model.UserDTO;
import pro.datnt.manabie.task.service.UserService;


@RestController
@RequiredArgsConstructor
@RequestMapping("/user")
@Slf4j
public class UserController {
    private final UserService userService;

    @PostMapping("/login")
    public ResponseEntity<?> login(@RequestBody UserDTO user) {

        return userService.loginUser(user)
                .map((token) -> {
                    LoginResponseDTO response = new LoginResponseDTO();
                    response.setToken(token);
                    return ResponseEntity.ok(response);
                })
                .orElse(ResponseEntity.badRequest().build());

    }

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody UserDTO user) {
        try {
            userService.createUser(user);
            return ResponseEntity.accepted().build();
        } catch (DataIntegrityViolationException e) {
            log.warn(e.getMessage(), e);
            return ResponseEntity.badRequest().build();
        } catch (Exception e) {
            log.error(e.getMessage(), e);
            return ResponseEntity.internalServerError().build();
        }
    }

}
