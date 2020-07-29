from django.shortcuts import render
from django.http import HttpResponse
import requests

def home(request):
    response = requests.get('http://localhost:8080/tools')
    toolsJson = response.json() 
    return render(request, 'website/home.html', {'tools':toolsJson})

def search(request):
    return render(request, 'website/search.html', {'title': 'Search tool'})
