package com.example.server.rest.response

import com.example.domain.model.TaskEntity
import java.time.LocalDateTime

data class TaskDto(
    val id: String,
    val title: String,
    val description: String,
    val userId: String,
    var createdAt: LocalDateTime = LocalDateTime.now(),
    var updatedAt: LocalDateTime = LocalDateTime.now(),
)

internal fun TaskEntity.toDisplayModel() = TaskDto(
    id = id,
    title = title,
    description = description,
    userId = userId,
    createdAt = createdAt,
    updatedAt = updatedAt,
)