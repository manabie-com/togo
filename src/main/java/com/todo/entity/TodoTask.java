package com.todo.entity;

import com.fasterxml.jackson.annotation.JsonBackReference;
import com.todo.common.LocalDateConverter;
import lombok.*;

import javax.persistence.*;
import java.sql.Date;
import java.time.LocalDate;
import java.util.Optional;


@Entity(name = "todo_task")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class TodoTask {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", columnDefinition = "BIGINT")
    private Long id;

    @Column(name = "content", columnDefinition = "VARCHAR(225)")
    private String content;

    @Column(name = "create_time", columnDefinition = "DATETIME")
    private Date createTime = new LocalDateConverter().convertToDatabaseColumn(LocalDate.now());

    @Column(name = "complete_time", columnDefinition = "DATETIME")
    private Date completeTime;

    @Column(name = "is_complete", columnDefinition = "BIT")
    private Boolean isComplete = false;

    @ManyToOne
    @JoinColumn(name = "app_account_id", referencedColumnName = "id", columnDefinition = "BIGINT")
    private AppAccount appAccount;

    public TodoTask(String content) {
        this.content = content;
    }
}
