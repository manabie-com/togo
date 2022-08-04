package com.example.server.infrastructure.task

import com.example.domain.model.TaskEntity
import com.example.domain.usecase.task.TaskDataSource
import org.springframework.data.r2dbc.repository.Modifying
import org.springframework.stereotype.Component
import org.springframework.transaction.annotation.Transactional

@Component
@Transactional
class TaskDS(
    private val taskRepo: TaskRepo,
) : TaskDataSource {

    @Modifying
    override suspend fun createTask(
        taskEntity: TaskEntity,
    ): TaskEntity {
        return taskRepo.save(
            Task.fromDomainModel(taskEntity)
        ).toDomainModel()
    }

    override suspend fun getNumberOfTaskOnToDay(userId: String): Long {
        return taskRepo.getTaskInDayByUserId(userId)
    }

}