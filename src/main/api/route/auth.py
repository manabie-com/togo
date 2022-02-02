from flask import Blueprint

auth_route = Blueprint('auth', __name__, url_prefix='/auth')


@auth_route.route("/")
def basic(**kwargs):
    return {"message": "Welcome to auth route"}
