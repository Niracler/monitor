basepath := $(shell pwd)

network:
	docker network create -d overlay --attachable spider

spider-nginx:
	docker service create --name spider_nginx \
--constraint node.role==manager \
-p 6379:6379 -p 27017:27017 -p 5000:5000 -p 6801-6804:6801-6804 \
--mount type=bind,source=$(basepath)/nginx/nginx-base.conf,target=/etc/nginx/nginx.conf \
--mount type=bind,source=$(basepath)/nginx/nginx-stream-proxy.conf,target=/etc/nginx/stream.conf.d/nginx-stream-proxy.conf \
--mount type=bind,source=$(basepath)/nginx/nginx-http-proxy.conf,target=/etc/nginx/conf.d/default.conf \
--network spider nginx:alpine

spider:
	docker stack deploy -c docker-compose.yml spider

scrapyd-image:
	docker build -t "plrom.niracler.com:5009/scrapyd" scrapyd
	docker push "plrom.niracler.com:5009/scrapyd"

webui-image:
	docker build -t "plrom.niracler.com:5009/webui" webui
	docker push "plrom.niracler.com:5009/webui"


