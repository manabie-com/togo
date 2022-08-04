package com.example.domain.exception

sealed class Error(
    open val code: Int,
    override val message: String,
    override val cause: Throwable?,
) : Throwable(message, cause) {

    companion object {
        const val UNEXPECTED_ERROR_STATUS_CODE = 500_000
        const val NOT_FOUND_ERROR_STATUS_CODE = 404_000
        const val OUT_OF_LIMIT_TASK_CODE = 400_001
        const val USER_IS_NOT_EXIT = 400_002
    }
    class UnexpectedError(
        cause: Throwable?,
    ) : Error(
        code = UNEXPECTED_ERROR_STATUS_CODE,
        message = "An unexpected error occurred",
        cause = cause,
    )

    class NotFoundError(
        cause: Throwable?,
    ) : Error(
        code = NOT_FOUND_ERROR_STATUS_CODE,
        message = "Not found",
        cause = cause,
    )

    object OutOfLimitTask : Error(
        code = OUT_OF_LIMIT_TASK_CODE,
        message = "out of limit task",
        cause = null,
    )

    object UserIsNotExit : Error(
        code = USER_IS_NOT_EXIT,
        message = "user is not exit",
        cause = null,
    )
}