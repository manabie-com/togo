package com.example.server.mapping

import com.example.domain.model.TaskEntity
import com.example.server.display.TaskDto

internal fun TaskEntity.toDisplayModel() = TaskDto(
    id = id,
    title = title,
    description = description,
    userId = userId,
    createdAt = createdAt,
    updatedAt = updatedAt,
)
