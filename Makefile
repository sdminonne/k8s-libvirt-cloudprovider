update_swagger:
	#For the moment update to the latest... Crossing fingers
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

validate_swagger:
	swagger validate api/swagger.json

generate_server: validate_swagger
	#go get -u -f ./...
	swagger generate server -f api/swagger.json -A k8s-libvirt-cloudprovider
