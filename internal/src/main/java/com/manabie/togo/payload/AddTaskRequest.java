package com.manabie.togo.payload;

import javax.validation.constraints.NotBlank;
import lombok.Data;


/**
 * The request of "addTask" handler: content
 * @author mupmup
 */
@Data
public class AddTaskRequest {
    
    @NotBlank
    private String content;
}
