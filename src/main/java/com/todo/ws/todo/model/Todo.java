package com.todo.ws.todo.model;

import com.todo.ws.todo.application.dto.TodoDTO;

import javax.persistence.*;
import java.time.LocalDate;

@Entity
public class Todo {

    @Id
    @SequenceGenerator(name = "TODO_SEQ", sequenceName = "TODO_SEQ", allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "TODO_SEQ")
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

    public Todo(TodoDTO todoDTO, Long todoUserId, LocalDate dateCreated) {
        this.task = todoDTO.getTask();
        this.status = todoDTO.getStatus();
        this.todoUserId = todoUserId;
        this.dateCreated = dateCreated;
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
