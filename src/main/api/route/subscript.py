from flask import Blueprint

subscription_route = Blueprint('subscription-route', __name__, url_prefix='/subscription')


@subscription_route.route("/")
def basic(**kwargs):
    return {"message": "Welcome to subscription route"}
