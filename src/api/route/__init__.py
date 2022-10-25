from .auth import auth_route
from .subscript import subscription_route
from .task import task_route

routes = [
    auth_route,
    task_route,
    subscription_route
]
