from flask import Flask, request, jsonify
import json
import uuid
from datetime import datetime
from task_daily_dao_module import TaskDaily
from task_dao_module import Task
from my_config import TABLE_TASKS, TABLE_DAILY, MAXIMUM_TASK_DAILY


app = Flask(__name__)
task_dao = Task(TABLE_TASKS)
task_daily_dao = TaskDaily(TABLE_DAILY)

def validate_task_limit(username):
    count = 0
    now = datetime.now().strftime('%Y-%m-%d')  
    list_task_by_user = task_daily_dao.get_all_tasks_daily()
    for task in list_task_by_user:
        if(task.get('owner') == username):
            count = count + 1
            task_date = (datetime. strptime(task.get('created_at'), '%Y-%m-%d %H:%M:%S')).strftime('%Y-%m-%d')
            if(count >= MAXIMUM_TASK_DAILY and now == task_date):
                return False
    return True

@app.route("/daily-task",  methods=['GET', 'POST'])
def daily_task():
    if request.method == 'POST':
        # curl -X POST http://127.0.0.1:5000/daily-task \
        # -H 'Content-Type: application/json' \
        # -d '{"project_id":"project1", "task":"task_hihi","owner":"hoaipham", "deadline_at": "2021-12-19 15:36:22"}'

        data = request.get_json(force=True)
        data_json =  json.loads(json.dumps(data))
        task_id = str(uuid.uuid4())
        now = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        insert_task = {
            'id': task_id,
            'task_name': data_json.get('task'),
            'project_id': data_json.get('project_id'),
            'created_at': now,
            'updated_at': now
        }
        id = str(uuid.uuid4())
        insert_task_daily = {
            'id': id,
            'task_id': task_id,
            'task': data_json.get('task'),
            'owner': data_json.get('owner'),
            'project_id': data_json.get('project_id'),
            'created_at': now,
            'updated_at': now,
            'deadline_at': data_json.get('deadline_at')
        }
        if(validate_task_limit(data_json.get('owner')) == True):
            task_dao.insert_task(insert_task)
            task_daily_dao.insert_task_daily(insert_task_daily)
            return jsonify({
                "status":"success",
                "message": "insert success",
                "status_code": 200
            })
        else:
            return jsonify({
                "status": "failed",
                "message": "number task is invalid due to maximum task daily",
                "status_code": 500
            })

    else:
        # curl http://127.0.0.1:5000/daily-task

        all_tasks = task_daily_dao.get_all_tasks_daily()
        return jsonify([json.loads(json.dumps(user)) for user in all_tasks])
if __name__ == '__main__':
    app.run()
