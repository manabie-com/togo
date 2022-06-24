package com.interview.challenges.domain;

import java.time.LocalDate;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;

import org.springframework.format.annotation.DateTimeFormat;
import org.springframework.format.annotation.DateTimeFormat.ISO;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer;

import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Entity
@Table(name = "tasks")
@NoArgsConstructor
public class Task{
	@Id
	@NotBlank(message = "id must not be blank")
	private String id;
	@NotBlank
	@NotBlank(message = "content must not be blank")
	private String content;
	@NotBlank
	private String userId;
	
	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd")
	@JsonDeserialize(using = LocalDateDeserializer.class)
	@JsonSerialize(using = LocalDateSerializer.class)
	@NotNull
	private LocalDate createdDate;
	
	public Task(String id, String content, String userId) {
		super();
		this.id = id;
		this.content = content;
		this.userId = userId;
	}
	
	public Task(String id, String content,String userId, LocalDate createdDate) {
		super();
		this.id = id;
		this.content = content;
		this.userId = userId;
		this.createdDate = createdDate;
	}
	
}
