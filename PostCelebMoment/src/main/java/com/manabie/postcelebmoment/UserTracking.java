package com.manabie.postcelebmoment;

import java.time.LocalDateTime;

public class UserTracking {
	private String userId;
	private LocalDateTime lastPostDate;
	private int limitation;
	private int counter;

	public int getCounter() {
		return counter;
	}

	public void setCounter(int counter) {
		this.counter = counter;
	}

	public UserTracking() {
	}

	public UserTracking(String userId, LocalDateTime lastPostDate, int limitation, int counter) {
		super();
		this.userId = userId;
		this.lastPostDate = lastPostDate;
		this.limitation = limitation;
	}

	public String getUserId() {
		return userId;
	}

	public void setUserId(String userId) {
		this.userId = userId;
	}

	public LocalDateTime getLastPostDate() {
		return lastPostDate;
	}

	public void setLastPostDate(LocalDateTime lastPostDate) {
		this.lastPostDate = lastPostDate;
	}

	public int getLimitation() {
		return limitation;
	}

	public void setLimitation(int limitation) {
		this.limitation = limitation;
	}
	
}