# Решение тасок курса https://education.nixsolutions.com/course/view.php?id=12

### run migration

```
yes | migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' down
migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up
```

### redis

Запуск redis cli

```
docker exec -it redis-token-store redis-cli
```

