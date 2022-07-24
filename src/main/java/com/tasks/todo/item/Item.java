package com.tasks.todo.item;

import org.springframework.data.annotation.Id;

public class Item {

    private final String userId;
    private final String taskName;
    private final String taskDescription;

    public Item(String userId, String taskName, String taskDescription) {
        this.userId = userId;
        this.taskName = taskName;
        this.taskDescription = taskDescription;
    }

    @Id
    public String getUserId() {
        return userId;
    }

    public String getTaskName() {
        return taskName;
    }

    public String getTaskDescription() {
        return taskDescription;
    }

    public Item updateWith(Item item) {
        return new Item(this.userId, item.taskName, item.taskDescription);
    }

}
