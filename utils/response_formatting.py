def get_format_message(message, http_status_code) -> tuple:
    return (
        {"message": message},
        http_status_code,
    )
