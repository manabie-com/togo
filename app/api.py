from flask import request, jsonify, make_response
from functools import wraps
from werkzeug.security import generate_password_hash, check_password_hash

import datetime
import uuid
import jwt

from .app import db
from .models import Users, Todos


def init_routes(app):
    def token_required(f):
        @wraps(f)
        def decorator(*args, **kwargs):
            token = None

            if "x-access-token" in request.headers:
                token = request.headers["x-access-token"]

            if not token:
                return jsonify({"message": "missing token."})

            try:
                print(token)
                data = jwt.decode(token, app.config["SECRET_KEY"], algorithms=["HS256"])
                current_user = Users.query.filter_by(public_id=data["public_id"]).first()
            except Exception as e:
                print(e)
                return jsonify({"message": "invalid token."})

            return f(current_user, *args, **kwargs)

        return decorator

    @app.route("/register", methods=["GET", "POST"])
    def signup_user():
        data = request.get_json()

        hashed_password = generate_password_hash(data["password"], method="sha256")

        new_user = Users(public_id=str(uuid.uuid4()), name=data["name"], limit_per_day=data["limit_per_day"],
                         password=hashed_password)
        db.session.add(new_user)
        db.session.commit()

        return jsonify({"message": "registration success."})

    @app.route("/login", methods=["GET", "POST"])
    def login_user():
        auth = request.authorization

        if not auth or not auth.username or not auth.password:
            return make_response("Verification error.", 401, {"WWW.Authentication": "Basic realm: 'login required'"})

        user = Users.query.filter_by(name=auth.username).first()

        if user is None:
            return make_response("Verification error.", 401, {"WWW.Authentication": "Basic realm: 'login required'"})

        if check_password_hash(user.password, auth.password):
            token = jwt.encode({"public_id": user.public_id,
                                "exp": datetime.datetime.utcnow() + datetime.timedelta(minutes=30)},
                               app.config["SECRET_KEY"])
            return jsonify({"token": token})

        return make_response("Verification Error", 401, {"WWW.Authentication": "Basic realm: 'login required'"})

    @app.route("/todo", methods=["POST"])
    @token_required
    def create_todo(current_user):

        data = request.get_json()
        # retrieve tasks of current user added today
        current_no_task = Todos.query.filter_by(user_id=current_user.id, date_time=datetime.date.today()).count()

        # return message if limit has been reached
        if current_no_task >= current_user.limit_per_day:
            return make_response("User has reached maximum todos per day.", 400)

        new_task = Todos(task=data["todo"], user_id=current_user.id, date_time=datetime.date.today())
        db.session.add(new_task)
        db.session.commit()

        return jsonify({"message": "new task added."})

    @app.route("/todos", methods=["GET"])
    @token_required
    def get_todos(current_user):

        todos = Todos.query.filter_by(user_id=current_user.id).all()

        result = []

        for todo in todos:
            data = {"id": todo.id, "task": todo.task, "date_time": todo.date_time}
            result.append(data)

        return jsonify({"todos": result})

    @app.route("/todo/delete", methods=["POST"])
    @token_required
    def delete_todo(current_user):
        todo_id = request.get_json()["todo_id"]
        todo = Todos.query.get(todo_id)

        if not todo_id:
            return jsonify({"message": "todo does not exist."})

        db.session.delete(todo)
        db.session.commit()

        return jsonify({"message": "todo deleted."})

    @app.route("/user/delete", methods=["POST"])
    @token_required
    def delete_user(current_user):
        user = Users.query.get(current_user.id)

        if not user:
            return jsonify({"message": "user does not exist."})

        db.session.delete(user)
        db.session.commit()

        return jsonify({"message": "user deleted."})
