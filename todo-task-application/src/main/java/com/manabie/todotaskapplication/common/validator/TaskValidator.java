package com.manabie.todotaskapplication.common.validator;

import com.manabie.todotaskapplication.common.constant.TaskActionType;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import org.apache.logging.log4j.util.Strings;
import org.springframework.stereotype.Component;

import java.util.Objects;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Component
public class TaskValidator {

    public boolean validateTaskDto(TaskDto taskDto, TaskActionType actionType) {
        if (Objects.isNull(taskDto)) {
            return false;
        }
        if (Strings.isEmpty(taskDto.getName())) {
            return false;
        }
        return true;
    }

}
