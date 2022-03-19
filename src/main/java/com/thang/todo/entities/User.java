package com.thang.todo.entities;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "users")
public class User {
    @Id
    @GeneratedValue
    @Column(name = "id")
    private Long id;

    @Column(name = "username", nullable = false, unique = true)
    private String username;

    @Column(name = "password")
    private String password;

    @Column(name = "maximum_tasks")
    private Long maximumTasks;

    public User() {}

    public User(String username, String password, Long maximumTasks) {
        this.username = username;
        this.password = password;
        this.maximumTasks = maximumTasks;
    }

    public Long getId() {
        return this.id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getUsername() {
        return this.username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return this.password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public Long getMaximumTasks() {
        return this.maximumTasks;
    }

    public void setMaximumTasks(Long maximumTasks) {
        this.maximumTasks = maximumTasks;
    }
}
