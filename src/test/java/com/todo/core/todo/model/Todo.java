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
    private Long TodoUserId;

    public Todo() {
    }

    public Todo(Long id, String status, String task, Long todoUserId) {
        this.id = id;
        this.status = status;
        this.task = task;
        TodoUserId = todoUserId;
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
        return TodoUserId;
    }
}
