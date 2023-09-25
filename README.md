

## 项目介绍:
- 使用阿里云[alibaba-cloud-sdk-go](github.com/aliyun/alibaba-cloud-sdk-go/services/cr) 获取镜像列表并生成auth.json和images.json文件
- 使用阿里云开源镜像同步工具[image-syncer](https://github.com/AliyunContainerService/image-syncer) 进行镜像同步

## 限制
- 仅支持从阿里云ACR同步到华为SWR
- 不支持同步到华为SWR企业版
- 同步过程中不会自动创建华为SWR组织，需要用户在同步之前手动创建

## 使用方法

### 1.下载
在[release](https://github.com/luochangbin/images-migrate/releases) 页面可直接下载二进制和源码包
### 2.创建配置文件config.yaml
``` bash
access_key: "xxxx" #阿里云ak
secret_key: "xxxx" #阿里云sk
region_ali: "cn-hangzhou" # 阿里云区域 https://help.aliyun.com/document_detail/198107.html
user_ali: "xxx" #阿里云镜像仓库登录用户
passwd_ali: "xxx" #阿里云镜像仓库登录密码
region_hw: "cn-south-1" # 华为云区域  https://developer.huaweicloud.com/endpoint?SWR
user_hw: "xxx" #华为云镜像仓库登录用户
passwd_hw: "xxx" #华为云镜像仓库登录密码
ns_org_map:   #可指定命名空间和组织名称同步关系，不配置该字段则将镜像同步至与命名空间同名的组织名称下
  namespace1: organization1  #将阿里云namespace1下的镜像同步至华为云organization1下
  namespace2: organization2
```
### 3.在华为云SWR控制台创建组织
### 4.执行命令
```bash
chmod 755 images-migrate-linux-amd64
./images-migrate-linux-amd64 -config config.yaml
```
### 5.执行结果
```bash
INFO[2023-09-22 11:15:37] Failed to log to file, using default stderr  
INFO[2023-09-22 11:15:37] Executing analyzing image rule for registry.cn-hangzhou.aliyuncs.com/test-acr/nginx -> swr.cn-south-1.myhuaweicloud.com/test-acr/nginx... 
INFO[2023-09-22 11:15:38] Finish analyzing image rule for registry.cn-hangzhou.aliyuncs.com/test-acr/nginx -> swr.cn-south-1.myhuaweicloud.com/test-acr/nginx. Now 1/1 tasks have been processed. 
INFO[2023-09-22 11:15:38] Executing generating sync tasks from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:39] Finish generating sync tasks from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 2/2 tasks have been processed. 
INFO[2023-09-22 11:15:39] Executing synchronizing blob sha256:605c77e624ddb75e6110f997c58876baa13f8754486b461117934b24a9dc3a85(7.656kB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:39] Executing synchronizing blob sha256:a0bcbecc962ed2552e817f45127ffb3d14be31642ef3548997f58ae054deb5b2(1.395kB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:39] Executing synchronizing blob sha256:b4df32aa5a72e2a4316aad3414508ccd907d87b4ad177abd7cbd62fa4dab2a2f(666B) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:39] Executing synchronizing blob sha256:186b1aaa4aa6c480e92fbd982ee7c08037ef85114fbed73dbb62503f24c1dd7d(894B) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:39] Executing synchronizing blob sha256:589b7251471a3d5fe4daccdddfefa02bdc32ffcba0a6d6a2768bf2c401faf115(602B) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:40] Finish synchronizing blob sha256:b4df32aa5a72e2a4316aad3414508ccd907d87b4ad177abd7cbd62fa4dab2a2f(666B) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 3/9 tasks have been processed. 
INFO[2023-09-22 11:15:40] Executing synchronizing blob sha256:a9edb18cadd1336142d6567ebee31be2a03c0905eeefe26cb150de7b0fbc520b(25.35MB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:40] Finish synchronizing blob sha256:589b7251471a3d5fe4daccdddfefa02bdc32ffcba0a6d6a2768bf2c401faf115(602B) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 4/9 tasks have been processed. 
INFO[2023-09-22 11:15:40] Executing synchronizing blob sha256:a2abf6c4d29d43a4bf9fbb769f524d0fb36a2edab49819c1bf3e76f409f953ea(31.36MB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:41] Finish synchronizing blob sha256:186b1aaa4aa6c480e92fbd982ee7c08037ef85114fbed73dbb62503f24c1dd7d(894B) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 5/9 tasks have been processed. 
INFO[2023-09-22 11:15:41] Finish synchronizing blob sha256:605c77e624ddb75e6110f997c58876baa13f8754486b461117934b24a9dc3a85(7.656kB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 6/9 tasks have been processed. 
INFO[2023-09-22 11:15:41] Finish synchronizing blob sha256:a0bcbecc962ed2552e817f45127ffb3d14be31642ef3548997f58ae054deb5b2(1.395kB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 7/9 tasks have been processed. 
INFO[2023-09-22 11:15:46] Finish synchronizing blob sha256:a9edb18cadd1336142d6567ebee31be2a03c0905eeefe26cb150de7b0fbc520b(25.35MB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 8/9 tasks have been processed. 
INFO[2023-09-22 11:15:50] Finish synchronizing blob sha256:a2abf6c4d29d43a4bf9fbb769f524d0fb36a2edab49819c1bf3e76f409f953ea(31.36MB) from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest: start to sync manifest. Now 9/9 tasks have been processed. 
INFO[2023-09-22 11:15:50] Executing synchronizing manifest from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest... 
INFO[2023-09-22 11:15:50] Finish synchronizing manifest from registry.cn-hangzhou.aliyuncs.com/test-acr/nginx:latest to swr.cn-south-1.myhuaweicloud.com/test-acr/nginx:latest. Now 10/10 tasks have been processed. 
INFO[2023-09-22 11:15:50] Finished, 0 tasks failed, cost 13.228387431s. 
```
## 更多参数
```bash
-h  --help       使用说明，会打印出一些启动参数的当前默认值

    -config      设置用户提供的配置文件路径
    
    -namespaces  同步指定命名空间下的镜像，多个命名空间用","，默认同步所有

    -log         打印出来的log文件路径，默认打印到标准错误输出，如果将日志打印到文件将不会有命令行输出，此时需要通过cat对应的日志文件查看

    -proc        并发数，进行镜像同步的并发goroutine数量，默认为5

    -retries     失败同步任务的重试次数，默认为2，重试会在所有任务都被执行一遍之后开始，并且也会重新尝试对应次数生成失败任务的生成。一些偶尔出现的网络错误比如io timeout、TLS handshake timeout，都可以通过设置重试次数来减少失败的任务数量

    -force       同步已经存在的、被忽略的镜像，这个操作会更新已存在镜像的时间戳

```
