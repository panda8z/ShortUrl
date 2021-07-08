build-dev-docker:
	# go build  -o shorturl main.go
	# docker build -t shorturl:v0.0.1-alpha.0.dev.1  -f dockerfile .
	docker build -t panda8z/surl:latest -t panda8z/surl:v0.0.1 -f Dockerfile.yaml

run-docker:
	docker run -it --name ttest -p 8080:8080 test

ping: 
	curl -X GET http://localhost:8080/ping

admin-ping:
	curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'

surl-post:
	curl -X POST \
	  	http://localhost:8080/surl \
	  	-H 'content-type: application/json' \
	  	-d '{"url":"http://localhost:8080/admin/bar"}'

surl-get: # c3VybDox is encoded form http://localhost:8080/admin/bar
	curl -X POST \
	  	http://localhost:8080/surl \
	  	-H 'content-type: application/json' \
	  	-d '{"url":"c3VybDox"}'

run-surl-restful-server:
	go run ./main.go

run-surl-gq-server:
	go run ./graph/cmd/graphql-server.go


run-hello-rpc-server:
	go run ./rpc/cmd/hello/server/rpc-server.go

run-hello-rpc-client:
	go run ./rpc/cmd/hello/client/rpc-client.go


gen-surl-proto:
	protoc --go_out=. --twirp_out=. surl.proto

run-surl-rpc-server:
	go run ./rpc/cmd/surl/rpc-server.go

run-surl-rpc-short:
	go run ./rpc/cmd/surl/rpc-client-short.go

run-surl-rpc-origin:
	go run ./rpc/cmd/surl/rpc-client-origin.go

# 经测试 无法在同一个终端窗口实现三个服务的开启。
# run-all-server:  run-surl-restful-server   run-surl-gq-srver  run-surl-rpc-server
