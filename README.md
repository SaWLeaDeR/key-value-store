# key-value-store
Key-Value-Store Application with go


```
Key-Value service
```

## Development

Run
```
go run cmd/server/main.go -env=env in cmd/server
```

## Testing

```
Make test
```

## Build

```
Make build
```


## Swagger

for swagger UI
```
Open http://127.0.0.1:9234/static/swagger-ui/
```

## Using Docker to simplify development (optional)

Run command in root folder
```
docker-compose up 
```

## Metrics

```
Open http://localhost:9090/ for Prometheus
Open http://localhost:16686/search for Jaeger UI
```

## Vault

```
You can use vault as a secure secrets management but you should add _SECURE  to environment variables like
I will use this for scheduler job timer property 
FILE_WRITER_CRON_TIMER_SECONDS_SECURE=0
FILE_WRITER_CRON_TIMER_MINUTES_SECURE=10
 You can access vault interface at localhost:8300 
 Method Token 
 Token = myroot
```

