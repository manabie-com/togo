package com.example.server.rest.task

import com.example.domain.model.TaskEntity
import com.example.domain.usecase.task.CreateTask
import com.example.domain.util.getOrElse
import com.example.server.rest.request.TaskRequest
import com.example.server.rest.response.TaskRespone
import com.example.server.rest.response.toDisplayModel
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.*

@Component
class TaskHandler(
    private val createTask: CreateTask,
) {

    suspend fun createTask(request: ServerRequest): ServerResponse {
        val req = request.awaitBody<TaskRequest>()

        val taskEntity = createTask(
            CreateTask.CreateTaskParam(
                TaskEntity(
                    title =  req.title,
                    description = req.description,
                    userId = req.userId
                )
            )
        ).getOrElse { throw it }
        return ServerResponse.ok().bodyValueAndAwait(
            TaskRespone(
                data = taskEntity.toDisplayModel()
            )
        )

    }
}
