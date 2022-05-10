package com.todo.entity;

import com.fasterxml.jackson.annotation.JsonBackReference;
import lombok.*;

import javax.persistence.*;
import java.util.List;


@Entity(name = "app_account")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class AppAccount {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", columnDefinition = "BIGINT")
    private Long id;

    @Column(name = "username", columnDefinition = "VARCHAR(40)")
    private String username;

    @Column(name = "password", columnDefinition = "VARCHAR(255)")
    private String password;

    @Column(name = "enabled", columnDefinition = "BIT")
    private Boolean enabled = true;

    @Column(name = "user_task_limit", columnDefinition = "BIGINT")
    private Integer userTaskLimit = 10;

    @OneToMany(mappedBy = "appAccount", cascade = CascadeType.ALL)
    @JsonBackReference
    private List<TodoTask> todoTaskList;
}
