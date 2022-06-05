import uuid
from django.db import models

from django.contrib.auth.models import AbstractBaseUser, AbstractUser
from django.contrib.auth.models import PermissionsMixin
from django.contrib.auth.validators import UnicodeUsernameValidator
from django.core.validators import MinValueValidator
from django.utils.translation import gettext_lazy as _

from .CustomUserManager import CustomUserManager

STATUS = (
    ('PD', 'Pending'),
    ('DO', 'Do'),
    ('CP', 'Completed')
)

PRIORITY = (
    ('L', 'LOW'),
    ('N', 'Normal'),
    ('H', 'High'),
    ('U', 'Urgent'),
    ('I', 'Immediate')
)


class CustomUser(AbstractBaseUser, PermissionsMixin):
    username_validator = UnicodeUsernameValidator()
    username = models.CharField(
        _("username"),
        max_length=50,
        unique=True,
        help_text=_(
            "Required. 50 characters or fewer. Letters, digits and @/./+/-/_ only."
        ),
        validators=[username_validator],
        error_messages={
            "unique": _("A user with that username already exists."),
        },
    )
    number_todo_limit = models.IntegerField(validators=[MinValueValidator(1)], default=5)

    USERNAME_FIELD = 'username'
    REQUIRED_FIELDS = []

    objects = CustomUserManager()
    @property
    def is_staff(self):
        return True
    @property
    def is_superuser(self):
        return True

    def __str__(self):
        return self.username


class Todo(models.Model):
    title = models.CharField(max_length=150, null=False, blank=False)
    description = models.TextField(null=True, blank=True)
    date = models.DateTimeField(auto_now_add=True)
    status = models.CharField(choices=STATUS, default='DO', max_length=10)
    priority = models.CharField(choices=PRIORITY, default='N', max_length=10)
    tag = models.ForeignKey(CustomUser, on_delete = models.CASCADE, null=True, blank=True)

    def __str__(self):
        return self.title


class TodoList(models.Model):
    id_user = models.ForeignKey(CustomUser, on_delete = models.CASCADE, null=True, blank=True)
    id_todo = models.ForeignKey(Todo, on_delete = models.CASCADE, null=True, blank=True)
    date = models.DateTimeField(auto_now_add=True)

    class Meta:
        unique_together = ('id_user', 'id_todo')
