from flask import request, jsonify, make_response
from functools import wraps
from werkzeug.security import generate_password_hash, check_password_hash

import datetime
import uuid
import jwt

from app.app import app, db
from app.models import Users, Todos

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

    new_user = Users(public_id=str(uuid.uuid4()), name=data["name"], limit_per_day=data["limit_per_day"], password=hashed_password)
    db.session.add(new_user)
    db.session.commit()

    return jsonify({"message": "registration success."})

@app.route("/login", methods=["GET", "POST"])
def login_user():
    auth = request.authorization

    if not auth or not auth.username or not auth.password:
        return make_response("Verification error.", 401, {"WWW.Authentication": "Basic realm: 'login required'"})

    user = Users.query.filter_by(name=auth.username).first()

    if check_password_hash(user.password, auth.password):
        token = jwt.encode({"public_id": user.public_id, "exp": datetime.datetime.utcnow() + datetime.timedelta(minutes=30)}, app.config["SECRET_KEY"])
        return jsonify({"token": token})
    
    return make_response("Verification Error", 401, {"WWW.Authentication": "Basic realm: 'login required'"})

@app.route("/users", methods=["GET"])
def get_all_users():

    users = Users.query.all()

    result = []

    for user in users:
        user_data = {}
        user_data["public_id"] = user.public_id
        user_data["name"] = user.name
        user_data["password"] = user.password

        result.append(user_data)

    return jsonify({"users": result})

@app.route("/todo", methods=["POST", "GET"])
@token_required
def create_task(current_user):
   
    data = request.get_json()
    # retrieve tasks of current user added today
    current_no_task = Todos.query.filter_by(user_id=current_user.id, date_time=datetime.date.today()).count()
    
    # return message if limit has been reached
    if (current_no_task >= current_user.limit_per_day):
        return jsonify({"message": "user has reached maximum todos per day."})

    new_task = Todos(task=data["todo"], user_id=current_user.id, date_time=datetime.date.today())
    db.session.add(new_task)   
    db.session.commit()   

    return jsonify({"message" : "new task added."})

@app.route("/todos", methods=["POST", "GET"])
@token_required
def get_todos(current_user):
    todos = Todos.query.filter_by(user_id=current_user.id).all()

    result = []

    for todo in todos:
        data = {}
        data["task"] = todo.task
        data["date_time"] = todo.date_time
        result.append(data)

    return jsonify({"todos": result})

# @app.route("/authors/<author_id>", methods=["DELETE"])
# @token_required
# def delete_author(current_user, author_id):
#     author = Authors.query.filter_by(id=author_id, user_id=current_user.id).first()
    
#     if not author:
#         return jsonify({"message": "Author does not exist."})
    
#     db.session.delete(author)
#     db.session.commit()

#     return jsonify({"message": "Author deleted."})