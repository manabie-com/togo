package com.todo.ws.common.utils;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.util.StringUtils;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;


import javax.servlet.http.HttpServletRequest;

public class JWTUtils {

    @Value("${todo.security.secret}")
    private String secret;

    public Long getUserIdFromJwt(String jwt) {
        return Long.valueOf(((Integer) getClaim(jwt, "todoUserId")).toString());
    }

    public Object getClaim(String token, String claim) {
        Claims claims = Jwts.parser()
            .setSigningKey(secret)
            .parseClaimsJws(token)
            .getBody();

        return claims.get(claim);
    }

    public String getJwtFromRequest(HttpServletRequest request) {
        String bearerToken = request.getHeader("Authorization");

        if(StringUtils.hasText(bearerToken) && bearerToken.startsWith("Bearer")) {
            return bearerToken.substring(7);
        }
        return null;
    }

    public String getUsernameFromToken(String token) {

        Claims claims = Jwts.parser()
            .setSigningKey(secret)
            .parseClaimsJws(token)
            .getBody();

        return claims.getSubject();
    }
}
