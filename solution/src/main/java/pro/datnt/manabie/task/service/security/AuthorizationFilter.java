package pro.datnt.manabie.task.service.security;

import io.jsonwebtoken.Claims;
import lombok.RequiredArgsConstructor;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;
import org.springframework.util.ObjectUtils;
import org.springframework.util.StringUtils;
import pro.datnt.manabie.task.properties.AuthenticationProperties;
import pro.datnt.manabie.task.service.UserService;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.Collections;
import java.util.Objects;

public class AuthorizationFilter extends BasicAuthenticationFilter {
    private final AuthenticationProperties properties;
    private final JwtUtil jwtUtil;
    private final UserService userService;

    public AuthorizationFilter(AuthenticationManager authenticationManager, AuthenticationProperties properties, JwtUtil jwtUtil, UserService userService) {
        super(authenticationManager);
        this.properties = properties;
        this.jwtUtil = jwtUtil;
        this.userService = userService;
    }

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain) throws IOException, ServletException {
        String header = request.getHeader(properties.getAuthorizationHeader());
        if (ObjectUtils.isEmpty(header)) {
            super.doFilterInternal(request, response, chain);
            return;
        }
        Claims claims = jwtUtil.parseToken(header);
        if (Objects.nonNull(claims)) {
            userService.getUserId(claims.getSubject())
                .ifPresent(id -> {
                    UsernamePasswordAuthenticationToken authenticationToken = new UsernamePasswordAuthenticationToken(id, null, Collections.emptyList());
                    SecurityContextHolder.getContext().setAuthentication(authenticationToken);
                });

        }
        super.doFilterInternal(request, response, chain);
    }
}
