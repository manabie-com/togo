package com.example.domain.model

import java.time.LocalDateTime
import java.util.*

data class TaskEntity(
    val id: String = UUID.randomUUID().toString(),
    val title: String,
    val description: String,
    val userId: String,
    var createdAt: LocalDateTime = LocalDateTime.now(),
    var updatedAt: LocalDateTime = LocalDateTime.now(),
)
