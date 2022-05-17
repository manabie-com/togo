from django.contrib.auth.base_user import BaseUserManager
from django.utils.translation import gettext_lazy as _


class CustomUserManager(BaseUserManager):

    def create_user(self, username, password, **extra_fields):
        """
        Create and save a User with the given username, password and todo_max.
        """
        if not username:
            raise ValueError(_('The username must be set'))
        user = self.model(username=username, **extra_fields)
        user.set_password(password)
        user.save()
        return user

    def create_superuser(self, username, password, **extra_fields):
        """
        Create and save a SuperUser with the given email, password and default todo_max = 1000.
        """
        extra_fields.setdefault('todo_max', 1000)

        if extra_fields.get('todo_max') != 1000:
            raise ValueError(_('Superuser must have todo_max is 1000.'))
        return self.create_user(username, password, **extra_fields)