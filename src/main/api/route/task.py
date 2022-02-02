from flask import Blueprint

task_route = Blueprint('task-route', __name__, url_prefix='/task')


@task_route.route("/")
def basic(**kwargs):
    return {"message": "Welcome to task route"}
