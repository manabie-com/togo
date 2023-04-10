exports.IN_SUFFICIENT_QUOTA = -1
exports.TOKEN_EXPIRED_MSG = 'TokenExpiredError'
exports.MAX_NUMBER_TASK_CREATED = 3
exports.ONE_DAY_TS = 86400

exports.ERROR = {
    USER_NOT_FOUND: {
        status_code: 400,
        error_msg: 'USER_NOT_FOUND'
    },
    USER_EXISTED: {
        status_code: 400,
        error_msg: 'USER_EXISTED'
    },
    TOKEN_EXPIRED: {
        status_code: 400,
        error_msg: 'TOKEN_EXPIRED'
    },
    INVALID_INPUT: {
        status_code: 400,
        error_msg: 'INVALID_INPUT'
    },
    INVALID_CREDENTIAL: {
        status_code: 400,
        error_msg: 'INVALID_CREDENTIAL'
    },
    IN_SUFFICIENT_QUOTA: {
        status_code: 400,
        error_msg: 'IN_SUFFICIENT_QUOTA'
    },
    UNAUTHORIZED: {
        status_code: 401,
        error_msg: 'UNAUTHORIZED'
    }
}

exports.COMMON_ERROR_CODE = {
    400: true,
    401: true,
    403: true,
}