package com.example.domain.model

import kotlinx.datetime.Clock
import kotlinx.datetime.Instant
import java.time.LocalDateTime
import java.util.*

data class TaskEntity(
    val id: String = UUID.randomUUID().toString(),
    //val id: Long = 5,
    val title: String,
    val description: String,
    val userId: String,
    var createdAt: LocalDateTime = LocalDateTime.now(),
    var updatedAt: LocalDateTime = LocalDateTime.now(),
)
