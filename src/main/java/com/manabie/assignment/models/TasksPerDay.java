package com.manabie.assignment.models;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToOne;
import java.sql.Date;

@Entity
public class TasksPerDay {
    @Id
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    private Integer Id;
    @OneToOne
    @JoinColumn(name="user_id")
    private User user;
    private Date taskDate;
    private Integer taskLimit;

    public Integer getId() {
        return Id;
    }

    public void setId(Integer id) {
        Id = id;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public Date getTaskDate() {
        return taskDate;
    }

    public void setTaskDate(Date taskDate) {
        this.taskDate = taskDate;
    }

    public Integer getTaskLimit() {
        return taskLimit;
    }

    public void setTaskLimit(Integer taskLimit) {
        this.taskLimit = taskLimit;
    }
}
