package com.manabie.interview.task.model;

import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.persistence.*;

@EqualsAndHashCode
@Getter
@Setter
@ToString
@Entity
@Table(name = "tasks")
public class Task {

    @Id
    @SequenceGenerator(
            name = "task_seq",
            sequenceName = "task_seq",
            allocationSize = 1
    )
    @GeneratedValue(
            strategy = GenerationType.SEQUENCE,
            generator = "task_seq"
    )
    private Long taskId;
    private String userUid;
    private String createdDate;

    private String taskDescription;
    
    public Task() {
    }


    public Task(String userUid, String createdDate, String taskDescription) {
        this.userUid = userUid;
        this.createdDate = createdDate;
        this.taskDescription = taskDescription;
    }

    public Task(Long taskId, String userUid, String createdDate, String taskDescription) {
        this.taskId = taskId;
        this.userUid = userUid;
        this.createdDate = createdDate;
        this.taskDescription = taskDescription;
    }

}
