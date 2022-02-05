from flask import Blueprint
from .route import routes

API_PREFIX = "/api"
api = Blueprint('api', __name__, url_prefix=API_PREFIX)
for route in routes:
    api.register_blueprint(route)


@api.route("/")
def home(**kwargs):
    return {"Message": "Welcome to Manabie-togo. Luu Hoang Son"}
