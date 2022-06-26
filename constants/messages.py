class HTTPReponseMessage:
    NOT_ALLOWED = "This user does not have permission to execute this action"
    MISSING_FIELD = "Missing field maximum_task_per_day"
    EXCEED_MAXIMUM_ERROR = "The number of tasks has exceeded the limit of per day"
    DELETE_SUCCESSFULL = "Delete successfully"
    NOT_ALLOWED_VIEW = "This user does not have permission to get information of others"
    INVALID_MAXIMUM_TASK_FIELD = (
        "The new value should be greater or equal to the current value"
    )
    NOT_ALLOWED_UPDATE_FIELD = "Not allowed to update %s field"
