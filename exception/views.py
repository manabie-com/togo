from django.db import IntegrityError
from rest_framework import status
from rest_framework.exceptions import AuthenticationFailed, ValidationError
from rest_framework.response import Response


def custom_exception_handler(exc, context):
    detail = ''
    code = status.HTTP_400_BAD_REQUEST

    if isinstance(exc, AuthenticationFailed):
        detail = 'Authentication fail'
        code = exc.status_code
    if isinstance(exc, ValidationError):
        detail = str(exc.detail[0]) if isinstance(exc.detail, (list, tuple)) else str(exc.detail)
        code = exc.status_code
    if isinstance(exc, IntegrityError):
        detail = str(exc)

    data = {
        'error': detail
    }

    return Response(data, status=code)

