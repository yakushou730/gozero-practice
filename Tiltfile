# -*- mode: Python -*-
print("hello tilt")

# docker build 這邊還動不了 先暫時用 build 好的 local image
# 指令要下 kind load docker-image store-api:v1
# docker_build("store-api:v1", ".", dockerfile="store/api/Dockerfile")

k8s_yaml('store/api/k8s.yaml')
k8s_resource('store-api', port_forwards=8888)
