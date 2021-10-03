# Решение тасок курса https://education.nixsolutions.com/course/view.php?id=12

### run migration

```
yes | migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' down
migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up
```

