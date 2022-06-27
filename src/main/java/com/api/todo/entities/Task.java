package com.api.todo.entities;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.validation.Valid;
import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;

import com.api.todo.request.RequestTaskEntity;

import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@NoArgsConstructor
@Data
@Table(name = "task")
public class Task {
    @Id
    @Column(name = "id")
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;

    @NotEmpty
    @Size(max = 50, message = "Title should have maximum 50 characters")
    private String title;

    @NotEmpty
    @Size(max = 250, message = "Description should have maximum 250 characters")
    private String description;

    @NotNull(message="User id is not empty")
    @Column(name = "user_id")
    private long userId;

    @Column(name = "created_date")
    private Date createdDate;

    @Column(name = "updated_date")
    private Date updatedDate;

    public Task() {}
    public Task(@Valid RequestTaskEntity requestTaskEntity) {
        this.title = requestTaskEntity.getTitle();
        this.description = requestTaskEntity.getDescription();
        this.userId = requestTaskEntity.getUserId();
        this.createdDate = new Date();
        this.updatedDate = new Date();
    }
    public Task(String title, String description, long userId) {
    	this.title = title;
        this.description = description;
        this.userId = userId;
        this.createdDate = new Date();
        this.updatedDate = new Date();
	}
	public long getId() {
        return id;
    }
    public void setId(long id) {
        this.id = id;
    }
    public String getTitle() {
        return title;
    }
    public void setTitle(String title) {
        this.title = title;
    }
    public String getDescription() {
        return description;
    }
    public void setDescription(String description) {
        this.description = description;
    }
    public long getUserId() {
        return userId;
    }
    public void setUserId(long userId) {
        this.userId = userId;
    }
    public Date getCreatedDate() {
        return createdDate;
    }
    public void setCreatedDate(Date createdDate) {
        this.createdDate = createdDate;
    }
    public Date getUpdatedDate() {
        return updatedDate;
    }
    public void setUpdatedDate(Date updatedDate) {
        this.updatedDate = updatedDate;
    }

}