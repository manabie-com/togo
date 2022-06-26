from django.db import models
from django.contrib.auth.models import (
    AbstractBaseUser,
    PermissionsMixin,
    BaseUserManager,
)


class CustomAccountManager(BaseUserManager):
    def create_superuser(self, username, password, **other_fields):

        other_fields.setdefault("is_staff", True)
        other_fields.setdefault("is_superuser", True)
        other_fields.setdefault("is_active", True)

        if other_fields.get("is_staff") is not True:
            raise ValueError("Superuser must be assigned to is_staff=True.")
        if other_fields.get("is_superuser") is not True:
            raise ValueError("Superuser must be assigned to is_superuser=True.")

        return self.create_user(username, password, **other_fields)

    def create_user(self, username, password, **other_fields):

        if not username:
            raise ValueError("You must provide an username")
        if not password:
            raise ValueError("You must provide a password")
        user = self.model(username=username, password=password, **other_fields)
        user.set_password(password)
        user.save()
        return user


class BaseUser(AbstractBaseUser, PermissionsMixin):

    username = models.CharField(max_length=150, unique=True, blank=False)
    is_active = models.BooleanField(default=True)
    created_at = models.DateTimeField(auto_now_add=True)
    maximum_task_per_day = models.PositiveIntegerField(default=10)
    is_staff = models.BooleanField(default=False)
    is_active = models.BooleanField(default=True)

    objects = CustomAccountManager()

    class Meta:
        db_table = "baseusers"
        ordering = ("-created_at",)

    USERNAME_FIELD = "username"
