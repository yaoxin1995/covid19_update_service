from django.contrib import admin
from django.urls import path
from user.views import register

from . import views



#path('', views.home, name='blog-home')  运用 function view 的url


# as_view()将一个class base view 转换为一个真正的view
urlpatterns = [
    #path('', PostListView.as_view(), name='blog-home'),
    path('', views.home, name='blog-home'),
    path('register/',register,name='register'),
    path('subscription/topic/', views.topicCreation, name='post-create'),
    path('about/',views.about,name='blog-about')
    #int:pk :primary key of post 
  #  path('post/<int:pk>/', PostDetailView.as_view(), name='post-detail'),

  # 对应于 post_detail 中的该 <h2><a class="article-title" href="{%url 'post-detail' post.id %}">{{ post.title }}</a></h2>
#	path('post/new/', views.creatTopic, name='post-create'),
#	path('post/<int:pk>/update/', PostUpdateView.as_view(), name='post-update'),
#	path('post/<int:pk>/delete/', PostDeleteView.as_view(), name='post-delete'),



  #accounts/login/
]
