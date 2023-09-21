

## 项目介绍:
- 利用阿里云接口获取镜像列表并生成auth.json和images.json文件
- 使用阿里云开源镜像工具[image-syncer](https://github.com/AliyunContainerService/image-syncer) 进行镜像同步

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
```
### 3.执行命令
```bash
chmod 755 images-migrate-linux-amd64
./images-migrate-linux-amd64 -config config.yaml
```

## 更多参数
```bash
-h  --help       使用说明，会打印出一些启动参数的当前默认值

    --config     设置用户提供的配置文件路径

    --log        打印出来的log文件路径，默认打印到标准错误输出，如果将日志打印到文件将不会有命令行输出，此时需要通过cat对应的日志文件查看

    --proc       并发数，进行镜像同步的并发goroutine数量，默认为5

    --retries    失败同步任务的重试次数，默认为2，重试会在所有任务都被执行一遍之后开始，并且也会重新尝试对应次数生成失败任务的生成。一些偶尔出现的网络错误比如io timeout、TLS handshake timeout，都可以通过设置重试次数来减少失败的任务数量

    --force      同步已经存在的、被忽略的镜像，这个操作会更新已存在镜像的时间戳

```
