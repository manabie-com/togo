package com.antulev.togo.models;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class UserInfo{
	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	private long id;
	
	private String uid;
	
	private long taskLimit;
	
	private String firstName;
	
	private String lastName;

	public long getId() {
		return id;
	}

	public void setId(long id) {
		this.id = id;
	}

	public String getUid() {
		return uid;
	}

	public void setUid(String uid) {
		this.uid = uid;
	}

	public long getTaskLimit() {
		return taskLimit;
	}

	public void setTaskLimit(long taskLimit) {
		this.taskLimit = taskLimit;
	}

	public String getFirstName() {
		return firstName;
	}

	public void setFirstName(String firstName) {
		this.firstName = firstName;
	}

	public String getLastName() {
		return lastName;
	}

	public void setLastName(String lastName) {
		this.lastName = lastName;
	}
	
}
