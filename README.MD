# Решение тасок курса https://education.nixsolutions.com/course/view.php?id=12

## Требования 
- ✅ CRUD для постов и комментов. Ответ от АПИ должен реализовать 2 формата представления данных: XML, JSON.
- gorm
- Перепиши реализованное АПИ с использованием echo фреймворка
- Swagger specification
- ✅ JWT
- OAuth 2.0
- Тестирование

### run migration

```
yes | migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' down
migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up
```

## Окружения для разработки

```
docker-compose up
```

### redis

Запуск redis cli

```
docker exec -it redis-token-store redis-cli
```

