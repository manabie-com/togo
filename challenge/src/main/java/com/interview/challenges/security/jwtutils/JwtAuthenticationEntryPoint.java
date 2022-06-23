package com.interview.challenges.security.jwtutils;

import java.io.IOException;
import java.io.PrintWriter;
import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.HashMap;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.json.JSONObject;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.json.JsonMapper;

@Component
public class JwtAuthenticationEntryPoint implements AuthenticationEntryPoint, Serializable {

	@Override
	public void commence(HttpServletRequest request, HttpServletResponse response,
			AuthenticationException authException) throws IOException, ServletException {
		JSONObject jsonObject = new JSONObject();
		jsonObject.put("timestamp", LocalDateTime.now());
		jsonObject.put("status", HttpServletResponse.SC_UNAUTHORIZED);
		jsonObject.put("error", "unauthorized");
		response.setContentType("application/json");
		response.getWriter().print(jsonObject);
	}

}
