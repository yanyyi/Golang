# Docker容器相关命令
查看容器:  
sudo docker ps #查看正在运行的容器
sudo docker ps -a #查看所有容器

创建容器:  
sudo docker run 参数  
参数说明  
· -i :保持容器运行  
· -t :为容器重新分配一个伪输入终端,通常与-i一起使用  
· -d :以守护(后台)模式运行容器。创建一个容器在后台运行,需要使用docker exec 进入容器。退出后,容器不会关闭  
· -it创建的一般称为交互式容器 -id创建的容器一般称为守护式容器
· --name :为创建的容器命名


进入容器:
sudo docker 参数 # 退出容器,容器不会关闭  
例子:sudo docker exec -it c2 /bin/bash

启动容器:  
sudo docker start 容器名称

停止容器:  
sudo docker stop 容器名称

删除容器:  
sudo docker rm 容器名称   #容器不能处于运行状态

查看容器信息:  
sudo docker inspect 容器名称