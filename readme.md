## development step

Reference: [go-zero](https://github.com/zeromicro/go-zero)

### prerequisite
1. install goctl
2. install kind as local k8s environment
   - `kind create cluster`
   - `kubectx kind-kind`

### create store api service
1. create API file (store.api)
```shell
mkdir store && cd store
goctl api -o store.api
```
2. generate corresponding files
```shell
goctl api go -api store.api -dir api 
```
3. run server
```shell
go run api/store.go -f api/etc/store-api.yaml
```
4. containerize store api service
```shell
cd api
goctl docker -go store.go -port 8888

# 回到 gozero-practice
docker build -t store-api:v1 -f store/api/Dockerfile .
docker run -p 8888:8888 store-api:v1
```
5. 加入 Tiltfile [reference](https://docs.tilt.dev/example_go.html)
- 設定 k8s.yaml (deployment)

> 因為 Tiltfile build image 現在不太會用 先跳過這部分 直接用 build 好的 local image
> 
> kind load docker-image store-api:v1
