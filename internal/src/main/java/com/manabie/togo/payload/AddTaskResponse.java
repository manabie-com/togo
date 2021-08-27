package com.manabie.togo.payload;

import com.manabie.togo.model.Task;
import java.util.List;
import lombok.Data;

/**
 * The response of "addTask" handler
 * @author mupmup
 */
@Data
public class AddTaskResponse {
    
    private List<Task> data;
}
