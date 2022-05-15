package com.manabie.todo.mapper;

import com.manabie.todo.domain.Task;
import com.manabie.todo.entity.TaskEntity;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

@Mapper
public interface TaskMapper {
  TaskMapper INSTANCE = Mappers.getMapper(TaskMapper.class);

  Task toDto(TaskEntity taskEntity);
}
