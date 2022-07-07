package com.todo.ws.core.auth.service;

import java.util.Date;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Component;

import io.jsonwebtoken.ExpiredJwtException;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.SignatureException;

@Component
public class TokenProviderService {

	   private static final Logger logger = LoggerFactory.getLogger(TokenProviderService.class);

	    @Value("${todo.security.secret}")
	    private String secret;

	    @Value("${todo.security.expiry}")
	    private int expiry;


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
