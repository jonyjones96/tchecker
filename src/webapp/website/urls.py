
from django.urls import path
from . import views

urlpatterns = [
    path('', views.home, name='website-home'),
    path('search/', views.search, name='website-search'),
]
