package com.todo.ws.core.todo.dto;

public class TodoDto {

    private String task;
    private String status = "not-completed";
    private Long todoUserId;

    public String getTask() {
        return task;
    }

    public void setTask(String task) {
        this.task = task;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Long getTodoUserId() {
        return todoUserId;
    }

    public void setTodoUserId(Long todoUserId) {
        this.todoUserId = todoUserId;
    }
}
