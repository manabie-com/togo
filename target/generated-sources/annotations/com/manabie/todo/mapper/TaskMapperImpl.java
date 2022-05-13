package com.manabie.todo.mapper;

import com.manabie.todo.domain.Task;
import com.manabie.todo.domain.Task.TaskBuilder;
import com.manabie.todo.entity.TaskEntity;
import javax.annotation.processing.Generated;

@Generated(
    value = "org.mapstruct.ap.MappingProcessor",
    date = "2022-05-13T19:31:06+0700",
    comments = "version: 1.4.2.Final, compiler: javac, environment: Java 17.0.2 (Oracle Corporation)"
)
public class TaskMapperImpl implements TaskMapper {

    @Override
    public Task toDto(TaskEntity taskEntity) {
        if ( taskEntity == null ) {
            return null;
        }

        TaskBuilder task = Task.builder();

        task.id( taskEntity.getId() );
        task.userId( taskEntity.getUserId() );
        task.title( taskEntity.getTitle() );
        task.description( taskEntity.getDescription() );
        task.datetimeCreated( taskEntity.getDatetimeCreated() );
        task.datetimeEdited( taskEntity.getDatetimeEdited() );
        task.isDelete( taskEntity.getIsDelete() );

        return task.build();
    }
}
