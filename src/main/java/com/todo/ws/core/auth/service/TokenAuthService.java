package com.todo.ws.core.auth.service;

import java.io.IOException;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.filter.OncePerRequestFilter;

import com.todo.ws.common.utils.JWTUtils;

@Service
public class TokenAuthService extends OncePerRequestFilter{

		@Autowired
	    private TokenProviderService tokenProvider;
		@Autowired
	    private UserDetailsService userDetailsService;
		@Autowired
	    private JWTUtils jwtUtils;

//	    public TokenAuthenticationFilter(TokenProvider tokenProvider, UserDetailsService userDetailsService, JwtUtils jwtUtils) {
//	        this.tokenProvider = tokenProvider;
//	        this.userDetailsService = userDetailsService;
//	        this.jwtUtils = jwtUtils;
//	    }

	    @Override
	    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
	        throws ServletException, IOException {

	        try {
	            String jwt = jwtUtils.getJwtFromRequest(request);

	            if(StringUtils.hasText(jwt) && tokenProvider.validateToken(jwt)) {
	                String username = jwtUtils.getUsernameFromToken(jwt);

	                UserDetails userDetails = userDetailsService.loadUserByUsername(username);
	                UsernamePasswordAuthenticationToken authentication
	                    = new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());
	                authentication.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));

	                logger.info("Username from Token: " + username);

	                SecurityContextHolder.getContext().setAuthentication(authentication);

	            }
	        }catch(Exception e) {
	            logger.error("Could not set authentication in Security Context.", e);
	        }

	        filterChain.doFilter(request, response);
	    }
}
