from django.contrib.auth.base_user import BaseUserManager
from django.utils.translation import gettext_lazy as _


class CustomUserManager(BaseUserManager):

    def create_user(self, username, password, **extra_fields):
        """
        Create and save a User with the given username, password and number_todo_limit.
        """
        if not username:
            raise ValueError(_('The username must be set'))
        # extra_fields.setdefault("is_staff", False)
        # extra_fields.setdefault("is_superuser", False)
        user = self.model(username=username, **extra_fields)
        user.is_admin = True
        user.set_password(password)
        user.save()
        return user


    def create_superuser(self, username, password, **extra_fields):
        """
        Create and save a SuperUser with the given email, password and default number_todo_limit_max = 5.
        """
        # extra_fields.setdefault("is_staff", True)
        # extra_fields.setdefault("is_superuser", True)
        extra_fields.setdefault('number_todo_limit', 5)

        if extra_fields.get('number_todo_limit') != 5:
            raise ValueError(_('Superuser must have number_todo_limit is 1000.'))
        return self.create_user(username, password, **extra_fields)
