# Mebius
kubenetes operator最佳实践

## Why Mebius
Mebius 是一个用Kubernetes Operator的思想，去管理多个Kubernetes集群，其中包括机器(节点)管理、集群管理、集群变更、策略/安全配置和事件分发的统一管理方案
![](docs/arch.png "整体架构图")

* Backend 主要是负责一些通用业务逻辑+agent方案管理，外加Controller
* Gateway 负责多个对接per cluster agent通信，也可以对接不同的Backend
  * 采用 grpc-gateway + gin 组件模式，即可以将grpc接口restful，也可以基于gin去开发一些webhook接口
* Agent 主要是负责获取远端配置+执行的Worker


# How to Build & clean
```
make build
make install
# clean
make clean
```

# How to run a components
For Example (Gateway)
```
./mebius gateway \ 
  --log_dir=/somewhere/log \
  --gateway-port=8023 \
  --logtostderr=false \
  --v=10 
```
