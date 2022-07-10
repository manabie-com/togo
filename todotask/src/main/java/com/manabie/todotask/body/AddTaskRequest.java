package com.manabie.todotask.body;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import java.time.ZonedDateTime;
import java.util.Date;

@Data
public class AddTaskRequest {
    @JsonProperty("target_date")
    private ZonedDateTime targetDate;
    @JsonProperty("user_id")
    private Integer userId;
    @JsonProperty("task_name")
    private String taskName;
    @JsonProperty("task_description")
    private String taskDescription;
}
