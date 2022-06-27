# 制作docker镜像

### step 1
编译项目
```bash
make
```

### step 2
在make 之后， 编译之后的文件输出在 `./output` 目录下，这个目录就是我们需要的文件。

**历史问题, 在制作之前需要手动修改deployment目录脚本的docker镜像版本。**

**这个版本号，就行你需要制作的版本号**
```text
// deployment/docker_env_install.sh
已经标注，请手动修改
// deployment/startnode.sh
已经标注，请手动修改
```

### step 3
如果需改为 v5.2.1,制作成docker镜像命令如下：
```bash
DOCKER_TAG=v5.2.1 make matrixchain-docker-image-no-compile
```

如果修改为 v5.2.2,制作docker镜像命令如下：
```bash
DOCKER_TAG=v5.2.2 make matrixchain-docker-image-no-compile
```

# 推送镜像到dockerhub

下面命令需要填写的参数：

    -DOCKER_TAG 版本号
    -dockerhub userid
    -dockerhub password

推送命令如下:
```bash
DOCKER_TAG=v5.2.2 DOCKER_USER=<your_name> DOCKER_PASSWORD=<your_password> make push-to-dockerhub
``` 