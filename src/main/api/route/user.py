from flask import Blueprint

user_route = Blueprint('user-route', __name__, url_prefix='/user')


@user_route.route("/", methods=['GET'])
def basic(**kwargs):
    return {"message": "Welcome to user route"}
