from rest_framework import serializers
from .models import UserRecordTask, UserTaskAllow


class UserRecordTaskSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserRecordTask
        fields = ["TaskTitle", "TaskDescription", "TaskDay"]#"__all__"


class UserTaskAllowSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserTaskAllow
        fields = ["task_allow", "task_done", "start_task_time", "last_task_time"]#"__all__"

