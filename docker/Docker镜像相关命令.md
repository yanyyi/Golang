# Docker镜像相关命令
查看镜像:查看本地所有镜像:  
sudo docker images  
sudo docker images -q #查看所有的镜像id

搜索镜像:从网络中查找需要的镜像:  
sudo docker search 镜像名称

拉取镜像:从Docker仓库下载镜像到本地,镜像名称格式为 名称:版本号:
sudo docker pull 镜像名称(:版本号)  #不指定默认下载latest版本

删除镜像:  
sudo docker rmi 镜像id #删除指定本地镜像
sudo docker rmi \`sudo docker images -q\` 
