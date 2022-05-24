package com.todo.core.user.service;

import com.todo.core.user.application.dto.CustomUserDetails;
import io.jsonwebtoken.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
public class TokenProvider {
    private static final Logger logger = LoggerFactory.getLogger(TokenProvider.class);

    @Value("${app.security.secret}")
    private String secret;

    @Value("${app.security.expiry}")
    private int expiry;

    public String generateToken(Authentication authentication) {

        CustomUserDetails userDetails = (CustomUserDetails) authentication.getPrincipal();
        Date date = new Date();
        Date expiryDate = new Date(date.getTime() + expiry);

        return Jwts.builder()
            .setSubject(userDetails.getUsername())
            .setIssuedAt(date)
            .setExpiration(expiryDate)
            .claim("todoUserId", userDetails.getId())
            .signWith(SignatureAlgorithm.HS512, secret)
            .compact();

    }

    public boolean validateToken(String authToken) {
        try {
            Jwts.parser().setSigningKey(secret).parseClaimsJws(authToken);
            return true;
        }catch (ExpiredJwtException e) {
            logger.error("Expired JWT Exception");
        }catch (SignatureException e) {
            logger.error("Invalid JWT signature");
        }
        return false;
    }

}
