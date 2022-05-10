from datetime import datetime

# validate the schedule passed using the following criteria:
# start time is LATER than current time
# start time is EARLIER than end time 

def validSchedule(request):
    start_dt = datetime.strptime(request.data.get("start_time"), '%Y-%m-%d %H:%M:%S')
    end_dt = datetime.strptime(request.data.get("end_time"), '%Y-%m-%d %H:%M:%S')
    return start_dt < end_dt and datetime.now() < start_dt
