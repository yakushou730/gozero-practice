run/store:
	go run store/api/store.go -f store/api/etc/store-api.yaml
run/product:
	go run product/rpc/product.go -f product/rpc/etc/product.yaml
