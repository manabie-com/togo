package com.example.server.rest.task

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.MediaType
import org.springframework.web.reactive.function.server.coRouter

@Configuration
class TaskRouter {
    companion object {
        const val TASK_API_ROUTE = "/api/tasks"
    }

    @Bean
    fun taskRoutes(
        taskHandler: TaskHandler,
    ) = coRouter {
        (accept(MediaType.APPLICATION_JSON) and TASK_API_ROUTE).nest {
            POST("", taskHandler::createTask)
        }
    }
}