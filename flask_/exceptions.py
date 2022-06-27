class APIAuthError(Exception):
    status_code = 401
    error_msg = "Authentication failed"
