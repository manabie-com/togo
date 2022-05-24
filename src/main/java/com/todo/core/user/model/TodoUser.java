package com.todo.core.user.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
public class TodoUser {

    @Id
    @GeneratedValue
    private Long id;

    @Column(nullable = false)
    private String username;

    @Column(nullable = false)
    private String password;

    @Column(nullable = false)
    private Long todoLimit;

    // Required by JPA
    protected TodoUser() {
    }

    public TodoUser(String username, String password, Long todoLimit) {
        this.username = username;
        this.password = password;
        this.todoLimit = todoLimit;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public Long getTodoLimit() {
        return todoLimit;
    }

    public void setTodoLimit(Long todoLimit) {
        this.todoLimit = todoLimit;
    }
}
