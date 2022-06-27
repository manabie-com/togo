from flask import Blueprint, g, make_response
from flask_.auth import auth_required
from flask_.blueprints.task import controller
from flask_.error import handle_exception

_app = Blueprint("task", __name__)


@_app.post("")
@handle_exception
@auth_required()
def post_task():
    response = controller.post_task(g.user_id)

    return make_response(*response)
