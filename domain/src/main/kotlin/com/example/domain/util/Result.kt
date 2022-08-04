package com.example.domain.util

import com.example.domain.exception.Error

sealed class Result<out V : Any> {

    fun isSuccess(): Boolean = this is Success

    fun isFailure(): Boolean = this is Failure

    fun getOrNull(): V? = when (this) {
        is Success -> value
        is Failure -> null
    }

    fun get(): V {
        check(isSuccess()) { "Could not get value of Failure Result" }
        return (this as Success).value
    }

    fun exceptionOrNull(): Error? = when (this) {
        is Success -> null
        is Failure -> error
    }

    fun exception(): Error {
        check(isFailure()) { "Could not get exception of Success Result" }
        return (this as Failure).error
    }

    class Success<out V : Any> internal constructor(val value: V) : Result<V>() {

        override fun toString() = "[Success: $value]"

        override fun hashCode(): Int = value.hashCode()

        override fun equals(other: Any?): Boolean {
            if (this === other) return true
            return other is Success<*> && value == other.value
        }
    }

    class Failure internal constructor(val error: Error) : Result<Nothing>() {

        override fun toString() = "[Failure: $error]"

        override fun hashCode(): Int = error.hashCode()

        override fun equals(other: Any?): Boolean {
            if (this === other) return true
            return other is Failure && error == other.error
        }
    }

    companion object {

        fun <V : Any> success(value: V): Result<V> = Success(value)

        fun failure(error: Throwable): Result<Nothing> = when (error) {
            is Error -> Failure(error)
            else -> Failure(Error.UnexpectedError(error))
        }
    }
}

fun <V : Any> result(f: () -> V) = try {
    Result.success(f())
} catch (e: Throwable) {
    Result.failure(e)
}

suspend fun <V : Any> suspendableResult(f: suspend () -> V) = try {
    Result.success(f())
} catch (e: Throwable) {
    Result.failure(e)
}

fun <V : Any> Result<V>.getOrElse(onFailure: (error: Error) -> V): V {
    return when (this) {
        is Result.Success -> value
        is Result.Failure -> onFailure(error)
    }
}