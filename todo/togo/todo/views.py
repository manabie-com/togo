from django.http.response import JsonResponse
from django.shortcuts import render
from django.http import HttpResponse
from django.views.decorators.csrf import csrf_exempt
from django.core import serializers
import json

from todo.models import User, Task

@csrf_exempt
def login(request):
    if request.method == 'POST':
        data = json.loads(request.body)

        userid = data.get('user_id')
        password = data.get('password')

        user = authenticate(userid, password)
        if user:
            request.session['auth'] = True
            return HttpResponse('Successfully logged in!')

    return HttpResponse('Failed to login!')


def tasks(request):
    if not request.session.has_key('auth') or not request.session['auth']:
        return HttpResponse('You need to login first!')

    if request.method == 'GET':
        date = request.GET.get('created_date')
        tasks = list(Task.objects.all().filter(date_created = date))

        tasks_list = []
        for task in list(tasks):
            item = {
                'content' : task.content,
                'date_created' : task.date_created
            }
            tasks_list.append(item)

        return JsonResponse(tasks_list, safe=False)
 
    return HttpResponse('No data found.')


def authenticate(userid, password):
    try:
        return User.objects.get(username = userid, password = password)
    except:
        return None