from celery_.services import task_service
from flask_.response import HttpResponse


def post_task(user_id):
    service_response = task_service.post_task.apply_async(
        args=[user_id], queue="task"
    ).get()

    return HttpResponse(*service_response)
