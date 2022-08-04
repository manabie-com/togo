package com.example.server.display

data class TaskRequest(
    val title: String,
    val description: String,
    val userId: String,
)