# dapr-clock
Example Clock for dapr

## Run the Docker Compose Definition

```bash
docker-compose up -d --build
```

## Request

```bash
docker-compose exec dev curl -i http://web/now
HTTP/1.1 200 OK
Content-Length: 33
Content-Type: application/json
Date: Tue, 01 Dec 2020 16:53:40 GMT
Server: fasthttp
Traceparent: 00-44c0e32fe59bff7ebc52485c6c91f1d8-593a60098e46962b-00

{"hour":0,"minute":3,"second":16}
```

or

```bash
docker-compose exec dev curl -i http://clock:3500/v1.0/invoke/clock/method/now
HTTP/1.1 200 OK
Server: fasthttp
Date: Tue, 01 Dec 2020 16:55:47 GMT
Content-Type: application/json
Content-Length: 33
Traceparent: 00-c5d4e795a7a7acc4677fe4f8b6d48766-f15f7f3ad6e71b1e-00

{"hour":0,"minute":5,"second":23}
```

## Clean up

```bash
docker-compose clean down
```
