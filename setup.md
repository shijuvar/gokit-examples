### Start CockroachDB cluster

cockroach start --insecure \
--store=ordersdb-1 \
--host=localhost
--background

cockroach start \
--insecure \
--store=ordersdb-2 \
--host=localhost \
--port=26258 \
--http-port=8081 \
--join=localhost:26257
--background

cockroach start --insecure \
--store=ordersdb-3 \
--host=localhost \
--port=26259 \
--http-port=8082 \
--join=localhost:26257 \
--background

### Start Zipkin in Docker
docker run -d -p 9411:9411 openzipkin/zipkin
