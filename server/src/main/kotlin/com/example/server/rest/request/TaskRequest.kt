package com.example.server.rest.request

data class TaskRequest(
    val title: String,
    val description: String,
    val userId: String,
)