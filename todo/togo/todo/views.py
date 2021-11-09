from django.shortcuts import render
from django.http import HttpResponse
from django.views.decorators.csrf import csrf_exempt
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
    if not request.session['auth']:
        return HttpResponse('You need to login first!')

    return HttpResponse('Hello World!!')

def authenticate(userid, password):
    try:
        return User.objects.get(id=userid, password=password)
    except:
        return None
