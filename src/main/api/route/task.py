from flask import Blueprint
from ..middleware import credentials_validation

task_route = Blueprint('task-route', __name__, url_prefix='/task')


# @task_route.route("/")
# def basic(**kwargs):
#     return {"message": "Welcome to task route"}
#

@task_route.route("/", methods=['GET'])
@credentials_validation
def task_route_get_all_task(**kwargs):
    payload = kwargs.get("payload")
    return {"message": f"Got all task of {payload}"}
