from rest_framework.response import Response


def response_success(data, status):
    data_response = {
        "data": data,
        "status": "success"
    }

    return Response(data_response, status=status)


def response_fail(data, status):
    data_response = {
        "data": {
            "error": data
        },
        "status": "fail"
    }

    return Response(data_response, status=status)
