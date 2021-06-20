package pro.datnt.manabie.task.service.security;

import io.jsonwebtoken.Claims;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import pro.datnt.manabie.task.properties.AuthenticationProperties;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.*;

class JwtUtilTest {
    AuthenticationProperties authenticationProperties = new AuthenticationProperties();
    JwtUtil jwtUtil = new JwtUtil(authenticationProperties);

    @BeforeEach
    public void setup() {
        authenticationProperties.setKey("1234234");
    }

    @Test
    void generateToken() {
        String token = jwtUtil.generateToken("test");
        Claims claims = jwtUtil.parseToken(token);
        assertThat(claims.getSubject()).isEqualTo("test");
    }
}