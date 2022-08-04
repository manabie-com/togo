package com.example.server.display

abstract class Response<D, E> {
    abstract val statusCode: Int
    abstract val error: String?
    open val result: D? = null
    open val extra: E? = null
}

data class ErrorResponse(
    override val statusCode: Int,
    override val error: String,
) : Response<Nothing, Nothing>()