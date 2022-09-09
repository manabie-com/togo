package com.manabie.todotaskapplication.data.pojo.task;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.List;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@NoArgsConstructor
@AllArgsConstructor
@Data
public class ResponseListTaskDto implements Serializable {
    private static final long serialVersionUID = 3620997384988593541L;
    private int total;
    private List<TaskDto> data;
}
