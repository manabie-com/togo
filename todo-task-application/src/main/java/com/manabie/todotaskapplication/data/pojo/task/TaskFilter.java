package com.manabie.todotaskapplication.data.pojo.task;

import com.manabie.todotaskapplication.data.pojo.filter.Filter;
import lombok.Data;

import java.util.List;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Data
public class TaskFilter extends Filter {
    private List<UUID> id;
    private List<String> name;
}
