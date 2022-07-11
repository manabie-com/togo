package com.manabie.todotask.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import javax.persistence.*;
import java.time.ZonedDateTime;
import java.util.Date;

@Data
@Entity
@Table(name = "`tasks`")
public class Task {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "`id`")
    private Integer id;

    @JsonProperty("user_id")
    @Column(name = "`user_id`")
    private Integer userId;

    @JsonProperty("description")
    @Column(name = "`description`")
    private String description;

    @JsonProperty("name")
    @Column(name = "`name`")
    private String name;

    @JsonProperty("target_date")
    @Column(name = "`target_date`")
    private ZonedDateTime targetDate;
}
