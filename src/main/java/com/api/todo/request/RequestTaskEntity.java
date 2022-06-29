package com.api.todo.request;

import java.io.Serializable;

import lombok.Builder;

@Builder
public class RequestTaskEntity implements Serializable {
    private static final long serialVersionUID = 1L;

    private String title;
    private String description;
    private long userId;

    public RequestTaskEntity() {}
    public RequestTaskEntity(String title, String description, long userId) {
		this.title = title;
		this.description = description;
		this.userId = userId;
	}
	public String getTitle() {
        return title;
    }
    public void setTitle(String title) {
        this.title = title;
    }
    public String getDescription() {
        return description;
    }
    public void setDescription(String description) {
        this.description = description;
    }
    public long getUserId() {
        return userId;
    }
    public void setUserId(long userId) {
        this.userId = userId;
    }
}
