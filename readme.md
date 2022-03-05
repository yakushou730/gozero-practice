## development step

Reference: [go-zero](https://github.com/zeromicro/go-zero)

### prerequisite
1. install goctl

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
