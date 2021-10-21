# Решение тасок курса [education.nixsolutions.com по Golang](https://education.nixsolutions.com/course/view.php?id=12)

## Требования

- ✅ настроить golint в Goland IDE (используем https://github.com/mgechev/revive,
  т.к. https://github.com/golang/go/issues/38968)
- ✅ CRUD для постов и комментов. Ответ от АПИ должен реализовать 2 формата представления данных: XML, JSON.
- ❌ GORM
- ✅ Перепиши реализованное АПИ с использованием echo фреймворка
- ✅ Swagger specification
- ✅ JWT
- ✅ OAuth 2.0 (Google)
- ✅ Тестирование

##Дополнительные задания (от себя)
- ✅ Добавить docker-compose для запуска окружения для локлаьной разработки
- ✅ Dockerfile
- Добавить конфигурирование приложения (например, https://github.com/spf13/viper#readme)
- Фронт (React)
- Настроить CI/CD

## Технологии

- Web framework https://github.com/labstack/echo
- Валидация https://github.com/go-playground/validator

- JWT https://jwt.io/
- Кеш для JWT-токенов https://redis.io/

- Миграции БД https://github.com/golang-migrate/migrate

- Документация https://github.com/swaggo/swag

- Mock-generator https://github.com/golang/mock
- Тестирование json ответа https://github.com/kinbiko/jsonassert

- Linter https://github.com/mgechev/revive


## Окружения для разработки

```bash
make docker-compose-dev-up
```

### run migration

```bash
yes | migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' down
migrate -path ./migrations/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up
```

### redis

```bash
#Запуск redis cli
docker exec -it redis-token-store redis-cli
```

## Прочее

```bash
# Описание Makefile
make help
```