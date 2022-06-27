package com.api.todo.entities;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.validation.constraints.NotEmpty;

import lombok.Builder;
import lombok.NoArgsConstructor;

@Entity
@NoArgsConstructor
@Table(name = "user")
@Builder
public class User {
    @Id
    @Column(name = "id")
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;

    @NotEmpty
    private String name;

    @Column(name="limit_tasks_per_day")
    private int limitTasksPerDay;

    public User(){}
    public User(long id, String name, int limitTasksPerDay) {
		this.id = id;
		this.name = name;
		this.limitTasksPerDay = limitTasksPerDay;
	}
    public User(String name, int limitTasksPerDay) {
		this.name = name;
		this.limitTasksPerDay = limitTasksPerDay;
	}

	public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public int getLimitTasksPerDay() {
        return limitTasksPerDay;
    }

    public void setLimitTasksPerDay(int limitTasksPerDay) {
        this.limitTasksPerDay = limitTasksPerDay;
    }
}