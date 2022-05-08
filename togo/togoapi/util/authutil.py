from ..models import User

def usernameExists(header_username):
    user = User.objects.filter(username=header_username)

    return user
