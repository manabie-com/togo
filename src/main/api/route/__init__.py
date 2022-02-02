from .auth import auth_route
from .subscript import subscription_route
from .task import task_route
from .user import user_route

routes = [
    auth_route,
    user_route,
    task_route,
    subscription_route
]
