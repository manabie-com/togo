from rest_framework import serializers

from todos.models import Task, SettingTask


class SettingTaskSerializer(serializers.ModelSerializer):
    user = serializers.PrimaryKeyRelatedField(read_only=True)

    class Meta:
        model = SettingTask
        fields = '__all__'

    def create(self, validated_data):

        task = SettingTask.objects.create(**validated_data)
        return task


class TaskSerializer(serializers.ModelSerializer):

    setting = SettingTaskSerializer(read_only=True)

    class Meta:
        model = Task
        fields = '__all__'

    def create(self, validated_data):
        setting = validated_data.get('setting')
        task_today = Task.objects.filter(setting=setting).count()
        if task_today >= setting.limit:
            raise serializers.ValidationError('Number of task is maximum today')

        task = Task.objects.create(**validated_data)
        return task
