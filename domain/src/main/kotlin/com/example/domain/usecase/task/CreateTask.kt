package com.example.domain.usecase.task

import com.example.domain.exception.Error
import com.example.domain.model.TaskEntity
import com.example.domain.usecase.UseCase
import com.example.domain.usecase.user.UserDateSource

class CreateTask(
    private val taskDataSource: TaskDataSource,
    private val userDataSource: UserDateSource,
) : UseCase<CreateTask.CreateTaskParam, TaskEntity>() {

    override suspend fun run(params: CreateTaskParam): TaskEntity {
        val user = userDataSource.getUser(params.task.userId) ?: throw Error.UserIsNotExist
        if (user.limitTask <= taskDataSource.getNumberOfTaskOnToDay(params.task.userId)) throw Error.OutOfLimitTask
        return taskDataSource.createTask(params.task)
    }

    data class CreateTaskParam(
        val task: TaskEntity
    ) : Params()
}