from celery import chain
from celery_.services import auth_service, task_service
from flask_.response import HttpResponse


def post_signin(signin_data):
    service_response = auth_service.post_signin.apply_async(
        args=[signin_data], queue="auth"
    ).get()

    return HttpResponse(*service_response)


def post_signup(signup_data):
    service_response = (
        chain(
            auth_service.post_signup.s(signup_data).set(queue="auth"),
            task_service.initialize_user_task.s().set(queue="task"),
        )
        .apply_async()
        .get()
    )
    return HttpResponse(*service_response)
