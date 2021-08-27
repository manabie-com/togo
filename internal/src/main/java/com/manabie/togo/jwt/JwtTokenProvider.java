package com.manabie.togo.jwt;

import java.util.Date;

import org.springframework.stereotype.Component;
import io.jsonwebtoken.*;
import com.manabie.togo.model.CustomUserDetails;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Utility to check JWT token
 *
 * @author mupmup
 */
@Component
public class JwtTokenProvider {

    private static final Logger log = LoggerFactory.getLogger(JwtTokenProvider.class);

    /**
     * Secret key of JWT
     */
    private final String JWT_SECRET = "manabietogo";

    /**
     * How long JWT will expired ?
     */
    private final long JWT_EXPIRATION = 604800000L;

    /**
     * Generate token from user details
     *
     * @param userDetails
     * @return JWT token
     */
    public String generateToken(CustomUserDetails userDetails) {
        // Lấy thông tin user
        Date now = new Date();
        Date expiryDate = new Date(now.getTime() + JWT_EXPIRATION);
        // Tạo chuỗi json web token từ id của user.
        return Jwts.builder()
                .setSubject(userDetails.getUsername())
                .setIssuedAt(now)
                .setExpiration(expiryDate)
                .signWith(SignatureAlgorithm.HS512, JWT_SECRET)
                .compact();
    }

    /**
     * Get user id from JWT token
     *
     * @param token
     * @return
     */
    public String getUserIdFromJWT(String token) {
        Claims claims = Jwts.parser()
                .setSigningKey(JWT_SECRET)
                .parseClaimsJws(token)
                .getBody();

        return claims.getSubject();
    }

    /**
     * Validate token
     *
     * @param authToken
     * @return
     */
    public boolean validateToken(String authToken) {
        try {
            Jwts.parser().setSigningKey(JWT_SECRET).parseClaimsJws(authToken);
            return true;
        } catch (MalformedJwtException ex) {
            log.error("Invalid JWT token");
        } catch (ExpiredJwtException ex) {
            log.error("Expired JWT token");
        } catch (UnsupportedJwtException ex) {
            log.error("Unsupported JWT token");
        } catch (IllegalArgumentException ex) {
            log.error("JWT claims string is empty.");
        }
        return false;
    }
}
