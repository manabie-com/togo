package com.thang.todo.payload;

import javax.validation.constraints.NotBlank;

public class TaskRequest {
    @NotBlank
    private String content;

    @NotBlank
    private String status;

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
}