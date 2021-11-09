import json
from datetime import datetime

from django.http import HttpResponse
from django.http.response import JsonResponse
from django.views.decorators.csrf import csrf_exempt

from todo.models import User, Task

@csrf_exempt
def login(request):
    if request.method == 'POST':
        data = json.loads(request.body)

        username = data.get('user_id')
        password = data.get('password')
        #returns authenticated user instance if username and password matched
        user = authenticate(username, password)
        #set session variables
        if user:
            request.session['auth'] = True
            request.session['user'] = user.id
            return HttpResponse('Successfully logged in!')

    return HttpResponse('Failed to login!', status=401)


@csrf_exempt
def tasks(request):
    #check if user is authenticated
    if request.session.is_empty():
        return HttpResponse('You need to login first!', status=401)
    
    auth_user_id = request.session['user']
    auth_user = User.objects.get(id=auth_user_id)
    #handles listing of tasks
    if request.method == 'GET':
        date = request.GET.get('created_date')
        date = datetime.strptime(date, '%Y-%m-%d').date()

        tasks = list(Task.objects.all()\
            .filter(date_created__year = date.year,
                    date_created__month = date.month,
                    date_created__day = date.day))

        tasks_list = []

        for task in list(tasks):
            item = {
                'content' : task.content,
                'date_created' : task.date_created
            }
            tasks_list.append(item)

        return JsonResponse(tasks_list, safe=False)
    #handles saving of tasks and checking the max limit
    if request.method == 'POST':
        data = json.loads(request.body)
        content = data.get('content')
        #check no of tasks for the day
        date_today = datetime.today()
        no_current_tasks = Task.objects.all()\
            .filter(user_id = auth_user,
                date_created__year = date_today.year,
                date_created__month = date_today.month,
                date_created__day = date_today.day).count()

        no_max_todo = auth_user.max_todo
        #save tasks if max is not reached
        if (no_current_tasks < no_max_todo):
            task = Task(content = content, user_id = auth_user)
            task.save()

            return HttpResponse('Successfully saved.')
        
        else:
            return HttpResponse('Max todo for the day already reached.', status=400)
 
    return HttpResponse(status=400)


def authenticate(userid, password):
    try:
        return User.objects.get(username = userid, password = password)
    except:
        return None
