from django.shortcuts import render
from django.http import HttpResponse
import requests
import os

def home(request):
    print("Getting host")
    host = os.environ['HOST_ENV']
    print("Host: {host}")
    response = requests.get('http://{}:8080/tools'.format(host))
    toolsJson = response.json() 
    return render(request, 'website/home.html', {'tools':toolsJson})

def search(request):
    return render(request, 'website/search.html', {'title': 'Search tool'})
