package com.manabie.todotaskapplication.data.pojo.task;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class TaskDto implements Serializable {
    private static final long serialVersionUID = 3620997384988593541L;
    private UUID id;
    private String name;
    private String content;
    private String userId;
}
