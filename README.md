

## 项目介绍:
- 利用阿里云接口获取镜像列表并生成auth.json和images.json文件
- 使用阿里云开源镜像工具[image-syncer](https://github.com/AliyunContainerService/image-syncer) 进行镜像同步

## 使用方法

### 1.编译成二进制
```bash
make build
```
### 2.创建配置文件config.yaml
``` bash
access_key: "xxxx" #阿里云ak
secret_key: "xxxx" #阿里云sk
region_ali: "cn-hangzhou" # 阿里云区域 https://help.aliyun.com/document_detail/198107.html
user_ali: "xxx" #阿里云镜像用户
passwd_ali: "xxx" #阿里云镜像密码
region_hw: "cn-south-1" # 华为云区域  https://developer.huaweicloud.com/endpoint?SWR
user_hw: "xxx" #华为云镜像用户
passwd_hw: "xxx" #华为云镜像密码
```
### 3.执行命令
```bash
./images-migrate -config config.yaml
```

## 更多参数
```bash
  -config string
        config file path (default "config.yaml")
  -force
        force update manifest whether the destination manifest exists
  -log string
        log file path (default in os.Stderr)
  -proc int
        numbers of working goroutines (default 5)
  -retries int
        times to retry failed task (default 2)

```
