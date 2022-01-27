package com.thang.todo.entities;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;
import java.util.Date;

@Entity
@Table(name = "tasks")
public class Task {
    @Id
    @GeneratedValue
    @Column(name = "id")
    private Long id;

    @Column(name = "content", nullable = false)
    private String content;

    @Column(name = "status")
    private String status;

    @Column(name = "created_date")
    private Date createdDate;

    @Column(name = "user_id")
    private Long userId;

    public Task() {}

    public Task(String content, String status, Date createdDate, Long userId) {
        this.content = content;
        this.status = status;
        this.createdDate = createdDate;
        this.userId = userId;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

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

    @Override
    public String toString() {
        return "Task{" +
                "id=" + id +
                ", content='" + content + '\'' +
                ", status='" + status + '\'' +
                ", createdDate=" + createdDate +
                ", userId=" + userId +
                '}';
    }
}
