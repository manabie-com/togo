from rest_framework.exceptions import APIException
from django.utils.translation import gettext_lazy as _


class Code200(APIException):
    status_code = 200
    default_detail = _("success")


class Code201(APIException):
    status_code = 201
    default_detail = _("new resource has been created")


class Code400(APIException):
    status_code = 400
    default_detail = _("parameters are wrong")


class Code404(APIException):
    status_code = 404
    default_detail = _("not found")


class Code401(APIException):
    status_code = 401
    default_detail = _("The request requires an user authentication")


class Code500(APIException):
    status_code = 500
    default_detail = _("the api developers error")


class OverTaskLimited(APIException):
    status_code = 500
    default_detail = _("user can't assign because over task limited")
