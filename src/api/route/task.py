from flask import Blueprint, request
import json
from ..middleware import credentials_validation
from ..controller.task import get_task_of, create_task, get_task_by, delete_task,update_task
from ..error.basic import HTTPError
from datetime import datetime

task_route = Blueprint('task-route', __name__, url_prefix='/task')


# @task_route.route("/")
# def basic(**kwargs):
#     return {"message": "Welcome to task route"}
#

@task_route.route("/", methods=['GET'])
@credentials_validation
def task_route_get_all_task(**kwargs):
    try:
        payload = kwargs.get("payload")
        user_id = payload.get("userId")
        tasks = get_task_of(user_id)
        return {"data": tasks}, 200
    except HTTPError as e:
        raise e
    except Exception as e:
        raise HTTPError(500, str(e))


@task_route.route("/", methods=['POST'])
@credentials_validation
def task_route_create_task(**kwargs):
    try:
        payload = kwargs.get("payload")
        user_id = payload.get("userId")
        body = json.loads(request.data)
        summary = body.get("summary", f"Default Title at {datetime.now().strftime('%Y-%m-%d')}")
        description = body.get("description", f"Default detail at {datetime.now().strftime('%Y-%m-%d')}")
        task_id = create_task(user_id, {
            "summary": summary,
            "description": description
        })
    except HTTPError as e:
        raise e
    except Exception as e:
        raise HTTPError(500, str(e))
    return {"message": f"Task created", "taskId": task_id}, 201


@task_route.route("/<task_id>", methods=["GET"])
@credentials_validation
def task_route_get_task_by_id(task_id, **kwargs):
    try:
        payload = kwargs.get("payload")
        user_id = payload.get("userId")
        task = get_task_by(task_id, user_id)
    except HTTPError as e:
        raise e
    except Exception as e:
        raise HTTPError(500, str(e))
    return {"data": task}, 200


@task_route.route("/<task_id>", methods=["PATCH"])
@credentials_validation
def task_route_update_task_by_id(task_id, **kwargs):
    try:
        payload = kwargs.get("payload")
        user_id = payload.get("userId")
        body = json.loads(request.data)
        task = update_task(task_id=task_id, user_id=user_id, task_data=body)
    except HTTPError as e:
        raise e
    except Exception as e:
        raise HTTPError(500, str(e))
    return {"data": task}, 200


@task_route.route("/<task_id>", methods=["DELETE"])
@credentials_validation
def task_route_delete_task_by_id(task_id, **kwargs):
    try:
        payload = kwargs.get("payload")
        user_id = payload.get("userId")
        result = delete_task(task_id, user_id)
        if result:
            return {"message": f"Task '{task_id}' deleted!"}, 200
    except HTTPError as e:
        raise e
    except Exception as e:
        raise HTTPError(500, str(e))

