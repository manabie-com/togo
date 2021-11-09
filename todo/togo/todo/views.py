from django.shortcuts import render
from django.http import HttpResponse

def login(request):
    return HttpResponse('Hello World!')

def tasks(request):
    return HttpResponse('Hello World!!')