package com.example.domain.usecase.task

import com.example.domain.model.TaskEntity

interface TaskDataSource {

    suspend fun createTask(
        taskEntity: TaskEntity
    ): TaskEntity

    suspend fun getNumberOfTaskOnToDay(
        userId: String
    ): Long
}