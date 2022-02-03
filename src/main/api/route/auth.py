from flask import Blueprint, request
import json
from ..controller.auth import create_user, validate_credential
from ..error import HTTPError

auth_route = Blueprint('auth', __name__, url_prefix='/auth')


@auth_route.route("/")
def basic(**kwargs):
    return {"message": "Welcome to auth route"}


@auth_route.route("/sign-up", methods=["POST"])
def auth_route_sign_up(**kwargs):
    try:
        body = json.loads(request.data)
        result = create_user(**body)
        return {"message": result}
    except Exception as e:
        raise HTTPError(500, str(e))


@auth_route.route("/sign-in", methods=["POST"])
def auth_route_sign_in(**kwargs):
    try:
        body = json.loads(request.data)
        email = body.get("email")
        password = body.get("password")
        result = validate_credential(email, password)
        if result != "":
            return {
                "token": result
            }
    except Exception as e:
        raise HTTPError(500, str(e))

    raise HTTPError(401, "Authentication Failed")
