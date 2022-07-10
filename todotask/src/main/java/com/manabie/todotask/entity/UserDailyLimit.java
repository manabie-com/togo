package com.manabie.todotask.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import javax.persistence.*;

@Entity
@Data
@Table(name = "`user_daily_limits`")
public class UserDailyLimit {
    @Id
    @JsonProperty("user_id")
    @Column(name="`user_id`")
    private Integer userId;

    @JsonProperty("daily_task_limit")
    @Column(name="`daily_task_limit`")
    private Integer dailyTaskLimit;
}
