package com.interview.challenges.utils;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.LinkedHashMap;

import org.springframework.http.HttpStatus;

import lombok.Data;

@Data
public class MessageBody extends LinkedHashMap<String, Object>{
	private LocalDate timestamp;
	private HttpStatus httpStatus;
	private String message;
	public MessageBody() {
		// TODO Auto-generated constructor stub
		this.timestamp = LocalDate.now();
	}
	public MessageBody(HttpStatus httpStatus, String message) {
		super();
		this.timestamp = LocalDate.now();
		this.httpStatus = httpStatus;
		this.message = message;
	}
	
	public void putAll() {
		put("timestamp", timestamp);
		put("status", HttpStatus.OK);
		put("message", message);
	}
}
