package com.todo.core.todo.model;

import com.todo.core.todo.application.dto.TodoDTO;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import java.time.LocalDate;

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

    @Column(nullable = false)
    private LocalDate dateCreated;

    public Todo() {
    }

    public Todo(String status, String task, Long todoUserId, LocalDate dateCreated) {
        this.status = status;
        this.task = task;
        this.todoUserId = todoUserId;
        this.dateCreated = dateCreated;
    }

    public Todo(TodoDTO todoDTO, Long todoUserId) {
        this.task = todoDTO.getTask();
        this.status = todoDTO.getStatus();
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

    public LocalDate getDateCreated() {
        return dateCreated;
    }
}
