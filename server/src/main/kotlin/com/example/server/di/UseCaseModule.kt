package com.example.server.di

import com.example.domain.usecase.task.CreateTask
import com.example.domain.usecase.task.TaskDataSource
import com.example.domain.usecase.user.UserDateSource
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

@Configuration
class UseCaseModule {
    @Bean
    fun createTaskUseCase(
        taskDataSource: TaskDataSource,
        userDateSource: UserDateSource,
    ): CreateTask = CreateTask(
        taskDataSource = taskDataSource,
        userDataSource = userDateSource,
    )
}