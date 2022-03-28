from flask import Flask, request

from database.models.queries import DBSession
from handler.task import HandlerTask
from datetime import datetime as dt

from database import init_databases, User, Calender, Assignment, Task

print("database loading...")
init_databases()
print("database loaded")

app = Flask(__name__)


@app.route("/")
def main():
    st = "Welcome to togo page"
    return st


@app.route("/api/demo", methods=["POST"])
def add_db_demo():
    if not request.json:
        raise
    try:
        db = DBSession()
        users = request.json["users"]
        for user in users:
            u = User(name=user.get("name"), ability=user.get("ability"))
            db.add(u)
        calendars = request.json["calenders"]
        for cal in calendars:
            c = Calender(date=dt.strptime(cal.get("date"), "%d/%m/%Y"))
            db.add(c)
        assignments = request.json["assignments"]
        for ass in assignments:
            a = Assignment(user_id=ass.get("user_id"), date=dt.strptime(ass.get("date"), "%d/%m/%Y"))
            db.add(a)
        tasks = request.json["tasks"]
        for task in tasks:
            t = Task(
                title=task.get("title"),
                description=task.get("description"),
                assignment_id=task.get("assignment_id")
            )
            db.add(t)
        return {"adding success": "added"}, 200
    except Exception as ex:
        print(str(ex))
        return {"internal occur": str(ex)}, 500


@app.route("/api/task/add", methods=['POST'])
def add_task():
    if not request.json:
        raise
    try:
        new_task = {
            'title': request.json["title"],
            'description': request.json.get("description", ""),
            'user_id': request.json.get("user_id"),
            'date': request.json.get("date")
        }
        db = DBSession()
        handler_task = HandlerTask(db)
        if handler_task.add_and_assign_task(new_task):
            return {"adding success": "added"}, 200
        return {"over ability": ""}, 200

    except Exception as ex:
        print(str(ex))
        return {"internal exception": str(ex)}, 500


if __name__ == '__main__':
    app.run(debug=True)
