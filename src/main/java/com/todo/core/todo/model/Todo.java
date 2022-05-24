package com.todo.core.todo.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
public class Todo {

    @Id
    @GeneratedValue
    private Long id;

    @Column(nullable = false)
    private String status;

    @Column(nullable = false)
    private String task;

    @Column(nullable = false)
    private Long todoUserId;

    public Todo() {
    }

    public Todo(String status, String task, Long todoUserId) {
        this.status = status;
        this.task = task;
        this.todoUserId = todoUserId;
    }

    public Long getId() {
        return id;
    }

    public String getStatus() {
        return status;
    }

    public String getTask() {
        return task;
    }

    public Long getTodoUserId() {
        return todoUserId;
    }
}
