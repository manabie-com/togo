from datetime import datetime
from django.utils import timezone
import pytz

from ..models import User

time_format = "%Y-%m-%d %H:%M:%S"

# validate the schedule passed using the following criteria:
# start time is LATER than current time
# start time is EARLIER than end time 
def validSchedule(request_start, request_end):
    if not request_start:
        return True

    start_dt = datetime.strptime(request_start, time_format).replace(tzinfo=pytz.utc)
    end_dt = datetime.strptime(request_end, time_format).replace(tzinfo=pytz.utc)
    return start_dt < end_dt and datetime.now(timezone.utc) <= start_dt

# compares daily limit to tasks sent the last 24 hours
# should return True if equal
def dailyLimitReached(user):
    return user.daily_limit == user.task_today
