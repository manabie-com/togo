package com.manabie.todo.configuration;

import com.manabie.todo.common.SecurityFilter;
import com.manabie.todo.service.TokenService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpHeaders;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.filter.OncePerRequestFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

import static com.manabie.todo.utils.Constants.JWT_TOKEN_INDEX;
import static com.manabie.todo.utils.Constants.PREFIX_TOKEN;

@AllArgsConstructor
@SecurityFilter
public class JwtFilterConfiguration extends OncePerRequestFilter {


  private final TokenService tokenService;

  @Override
  protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response,
                                  FilterChain filterChain) throws ServletException, IOException {

    var authorizationHeader = request.getHeader(HttpHeaders.AUTHORIZATION);
    if (authorizationHeader == null || !authorizationHeader.startsWith(PREFIX_TOKEN)) {
      doFilter(request, response, filterChain);
      return;
    }

    var tokenParts = authorizationHeader.split(PREFIX_TOKEN);
    var jwtToken = tokenParts[JWT_TOKEN_INDEX];
    var user = tokenService.validate(jwtToken);
    if (user.isEmpty()) {
      doFilter(request, response, filterChain);
      return;
    }

    var authentication = new UsernamePasswordAuthenticationToken(user.get(), null, null);
    SecurityContextHolder.getContext().setAuthentication(authentication);
    doFilter(request, response, filterChain);
  }
}
