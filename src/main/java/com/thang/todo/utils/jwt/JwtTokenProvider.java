package com.thang.todo.utils.jwt;

import java.util.Date;

import com.thang.todo.services.UserDetailsImpl;

import org.slf4j.*;
import org.springframework.stereotype.Component;

import io.jsonwebtoken.*;

@Component
public class JwtTokenProvider {
    Logger logger = LoggerFactory.getLogger(JwtTokenProvider.class);

    private final String JWT_SECRET = "secret";

    private final long JWT_EXPIRATION = 604800000L;

    public String generateToken(UserDetailsImpl userDetails) {
        Date now = new Date();
        Date expiryDate = new Date(now.getTime() + JWT_EXPIRATION);
        return Jwts.builder()
                   .setSubject(Long.toString(userDetails.getUser().getId()))
                   .setIssuedAt(now)
                   .setExpiration(expiryDate)
                   .signWith(SignatureAlgorithm.HS512, JWT_SECRET)
                   .compact();
    }

    public Long getUserIdFromJWT(String token) {
        Claims claims = Jwts.parser()
                            .setSigningKey(JWT_SECRET)
                            .parseClaimsJws(token)
                            .getBody();

        return Long.parseLong(claims.getSubject());
    }

    public boolean validateToken(String authToken) {
        try {
            Jwts.parser().setSigningKey(JWT_SECRET).parseClaimsJws(authToken);
            return true;
        } catch (MalformedJwtException ex) {
            logger.error("Invalid JWT token");
        } catch (ExpiredJwtException ex) {
            logger.error("Expired JWT token");
        } catch (UnsupportedJwtException ex) {
            logger.error("Unsupported JWT token");
        } catch (IllegalArgumentException ex) {
            logger.error("JWT claims string is empty.");
        }
        return false;
    }
}