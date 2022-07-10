package com.manabie.todotask.body;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import javax.validation.constraints.NotNull;
import java.time.ZonedDateTime;

@Data
public class AddTaskRequest {
    @JsonProperty("target_date")
    @NotNull(message = "Target date must not be null")
    private ZonedDateTime targetDate;

    @JsonProperty("user_id")
    @NotNull(message = "User id must not be null")
    private Integer userId;

    @JsonProperty("task_name")
    private String taskName;
    @JsonProperty("task_description")
    private String taskDescription;
}
