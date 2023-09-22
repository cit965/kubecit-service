# WLB965

### 项目介绍

在线成人技能课程平台后端服务，提供课程展示，讲师入驻，在线编程等功能，项目使用golang语言kratos框架进行开发。


## ent

go run -mod=mod entgo.io/ent/cmd/ent new Category

go generate ./ent


## 部署

### k8s yaml
k8s yaml 在 deploy 文件夹下,包含 deployment，service，ingress

### 打包镜像

```shell
docker buildx create --name mybuilder
docekr buildx use mybuilder
docker buildx build --platform linux/amd64 -t chaoyue/kubecit-service --push .
```

上传到 dockerhub 下 chaoyue/kubecit-service 


## api 地址 

服务启动后访问 http://localhost:8000/q/swagger-ui


## 链接 qa 数据库

需要有 qa 的kubeconfig ，然后执行

```shell
kubectl port-forward service/wlb965-mysql 3306:3306
make runqa
```

