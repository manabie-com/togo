package com.example.server.display

data class TaskRespone(
    val code: Int = 0,
    val message: String = "success",
    val data: TaskDto? = null,
)