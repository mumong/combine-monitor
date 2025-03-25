#!/bin/bash

# 配置区域 - 在这里修改镜像名称和标签
IMAGE_NAME="queryservice"    # 镜像名称
IMAGE_TAG="v1.0"            # 镜像标签
FULL_IMAGE="${IMAGE_NAME}:${IMAGE_TAG}"  # 完整镜像名称

# Step 1: 删除所有使用指定镜像启动的容器
containers=$(docker ps -a -q --filter ancestor="$FULL_IMAGE")

if [ -n "$containers" ]; then
    echo "以下容器使用了 $FULL_IMAGE 镜像，将被强制删除:"
    echo "$containers"
    docker rm -f $containers
else
    echo "没有找到使用 $FULL_IMAGE 的容器。"
fi

# Step 2: 删除指定镜像
echo "删除 $FULL_IMAGE 镜像..."
docker rmi -f "$FULL_IMAGE"
crictl rmi "docker.io/library/$FULL_IMAGE"
docker rmi $(docker images -f "dangling=true" -q)

# Step 3: 删除所有 <none> 镜像
echo "删除所有悬挂的 <none> 镜像..."
docker image prune -f

# Step 4: 删除所有停止的容器
echo "删除所有已停止的容器..."
docker container prune -f

# Step 5: 删除所有未使用的网络
echo "删除所有未使用的网络..."
docker network prune -f

# Step 6: 删除所有未使用的卷
echo "删除所有未使用的卷..."
docker volume prune -f

echo "Docker 清理完成。"
