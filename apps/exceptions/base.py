from rest_framework.views import exception_handler
from rest_framework.exceptions import ValidationError
from rest_framework.response import Response


def exceptions_handler(exc, context):
    response = exception_handler(exc, context)
    if response is not None:
        msg = ""
        if isinstance(exc, ValidationError):
            temp = response.data
            response.data = dict()
            for i in temp:
                try:
                    msg += " {} {}".format(i, str(temp[i][0]))
                except Exception as e:
                    for j in temp[i]:
                        msg += " {} {}".format(i, str(temp[i][j][0]))
        else:
            msg = response.data.pop('detail')
        response.data['status'] = exc.status_code
        response.data['message'] = msg
        response.data['data'] = None
    else:
        data = {'status': exc.status_code, 'message': str(exc), 'data': None}
        response = Response(data, status=exc.status_code, headers=None)
    return response


def make_success_response(data=None, status_code=None, headers=None, message='success'):
    data_response = {
        'status': status_code, 'message': message, 'data': data
    }
    return Response(data_response, status=status_code, headers=headers)
