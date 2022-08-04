package com.example.server.display

import java.time.LocalDateTime

data class TaskDto(
    val id: String,
    val title: String,
    val description: String,
    val userId: String,
    var createdAt: LocalDateTime = LocalDateTime.now(),
    var updatedAt: LocalDateTime = LocalDateTime.now(),
)