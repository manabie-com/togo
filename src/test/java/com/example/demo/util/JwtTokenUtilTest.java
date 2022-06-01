package com.example.demo.util;

import com.example.demo.model.User;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.platform.commons.util.StringUtils;

import java.util.Calendar;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import static com.example.demo.util.JwtTokenUtil.JWT_TOKEN_VALIDITY;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class JwtTokenUtilTest {

    private final String SECRET="test-secret";

    private User user;
    private JwtTokenUtil jwtTokenUtil;
    @BeforeEach
    public void init() {

        Calendar cal = Calendar.getInstance();
        cal.set(Calendar.YEAR, 2022);
        cal.set(Calendar.MONTH, Calendar.DECEMBER);
        cal.set(Calendar.DAY_OF_MONTH, 2);
        Date expirationDate = cal.getTime();
        user = new User("test", "password");
        jwtTokenUtil = new JwtTokenUtil();
    }

    @Test
    public void testGenerateJwt() {
        Map<String, Object> claims = new HashMap<>();
        String jwt = jwtTokenUtil.generateToken(user);
        String username = jwtTokenUtil.getUsernameFromToken(jwt);
        assertFalse(StringUtils.isBlank(jwt));
        assertEquals("test", username);
        assertTrue(jwtTokenUtil.validateToken(jwt, user));
    }
}
