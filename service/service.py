import os
from flask import Flask, request, Response
import json
from controllers.user_login_controller import LoginController
from controllers.user_register_controller import UserRegisterController
from controllers.task_controller import TaskController
from controllers.authentication_controller import Authentication

app = Flask(__name__)


@app.route('/user_register', methods=['POST'])
def user_register():
    data = request.json
    print(data, type(data))
    result = UserRegisterController().register(data)
    return Response(json.dumps(result).encode('utf8'), mimetype='application/json', status=result["status"])


@app.route('/login', methods=['POST'])
def login():
    data = request.json
    print("------------", data, type(data))
    result = LoginController().login(data=data)
    return Response(json.dumps(result).encode('utf8'), mimetype='application/json', status=result["status"])


@app.route('/task', methods=['POST'])
def create_task():
    data = request.json
    token = request.headers.get('Authorization')
    user_info = Authentication().validate_token(token=token)
    if not user_info:
        rsl = {'status': 401, "message": "Non-Authoritative Information!"}
        return Response(json.dumps(rsl).encode('utf8'), mimetype='application/json', status=rsl["status"])
    else:
        task = TaskController().create_task(data=data, user_info=user_info)
        return Response(json.dumps(task).encode('utf8'), mimetype='application/json', status=task["status"])


@app.route('/user', methods=['PATCH'])
def update_user():
    data = request.json
    user_name = data.get('user_name')
    password = data.get('password')
    print(user_name, password)
    rsl = {'status': 200, "message": "SUCCESS!"}
    return Response(json.dumps(rsl).encode('utf8'), mimetype='application/json')


@app.route('/task', methods=['GET'])
def get_task():
    token = request.headers.get('Authorization')
    user_info = Authentication().validate_token(token=token)
    if not user_info:
        rsl = {'status': 401, "message": "Non-Authoritative Information!"}
        return Response(json.dumps(rsl).encode('utf8'), mimetype='application/json', status=rsl["status"])
    else:
        task = TaskController().list_task(user_info=user_info)
        return Response(json.dumps(task).encode('utf8'), mimetype='application/json', status=task["status"])


@app.route('/ping', methods=['GET'])
def ping():
    return "ok"

