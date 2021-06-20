package pro.datnt.manabie.task.service.security;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Component;
import pro.datnt.manabie.task.properties.AuthenticationProperties;

import java.util.Date;

@RequiredArgsConstructor
@Component
public class JwtUtil {
    private final AuthenticationProperties properties;

    public String generateToken(String userName) {
        Date expirationTime = new Date(System.currentTimeMillis() + properties.getExpirationTime());
        Claims claims = Jwts.claims().setSubject(userName);
        return Jwts.builder()
                .setClaims(claims)
                .setExpiration(expirationTime)
                .signWith(SignatureAlgorithm.HS512, properties.getKey())
                .compact();
    }

    public Claims parseToken(String token) {
        return Jwts.parser()
                .setSigningKey(properties.getKey())
                .parseClaimsJws(token)
                .getBody();
    }

}
