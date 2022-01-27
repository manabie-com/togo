package com.thang.todo.payload;

import javax.validation.constraints.NotBlank;
import java.util.Date;

public class TaskRequest {
    @NotBlank
    private String content;

    @NotBlank
    private String status;

    private Date createdDate;

    private Long userId;

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Date getCreatedDate() {
        return createdDate;
    }

    public void setCreatedDate(Date createdDate) {
        this.createdDate = createdDate;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }
}