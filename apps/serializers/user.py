from rest_framework import serializers

from apps.exceptions.status_code import Code500, OverTaskLimited
from apps.models.models.detail import Detail


class DetailSerializer(serializers.ModelSerializer):
    schedule = serializers.CharField(read_only=True)
    taskmaster = serializers.CharField(read_only=True)
    task = serializers.IntegerField(required=True, write_only=True)
    user = serializers.IntegerField(required=True, write_only=True)
    date = serializers.DateField(required=True, write_only=True)
    limit = serializers.IntegerField(read_only=True)

    class Meta:
        model = Detail
        fields = ('schedule', 'taskmaster', 'task', 'user', 'date', 'limit')

    def create(self, validated_data):
        from apps.views.user import CreateDetail
        try:
            user = validated_data.pop('user')
            date_ = validated_data.pop('date')
            schedule_task = CreateDetail.get_schedule_task(user, date_)
            limit_number_task = schedule_task.limit
            current_number_task = len(schedule_task.details.all())
            if current_number_task > limit_number_task:
                raise OverTaskLimited

            detail = {
                'schedule': schedule_task,
                'taskmaster': self.context['request'].user,
                'task': validated_data.pop('task'),
            }
            return Detail.objects.create(**detail)
        except Exception as e:
            raise
