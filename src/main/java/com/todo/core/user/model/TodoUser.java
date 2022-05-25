package com.todo.core.user.model;

import com.todo.core.user.application.dto.UserRegistrationDTO;

import javax.persistence.*;

@Entity
public class TodoUser {

    @Id
    @SequenceGenerator(name = "USER_SEQ", sequenceName = "USER_SEQ", allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "USER_SEQ")
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

    public TodoUser(UserRegistrationDTO dto) {
        this.username = dto.getUsername();
        this.password = dto.getPassword();
        this.todoLimit = (long) dto.getTodoLimit();
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
