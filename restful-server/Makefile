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