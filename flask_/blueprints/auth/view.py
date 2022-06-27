from flask import Blueprint, make_response, request
from flask_.blueprints.auth import controller
from flask_.blueprints.auth.schema import SignInSchema, SignUpSchema
from flask_.error import handle_exception

_app = Blueprint("auth", __name__)


@_app.post("/signup")
@handle_exception
def post_signup():
    signup_data = SignUpSchema().load(request.json)
    response = controller.post_signup(signup_data)

    return make_response(*response)


@_app.post("/signin")
@handle_exception
def post_signin():
    signup_data = SignInSchema().load(request.json)
    response = controller.post_signin(signup_data)

    return make_response(*response)
