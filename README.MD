![Github CI/CD](https://img.shields.io/github/workflow/status/asavt7/nixedu-backend/Go)
![Go Report](https://goreportcard.com/badge/github.com/asavt7/nixedu-backend)
![Repository Top Language](https://img.shields.io/github/languages/top/asavt7/nixedu-backend)
![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/asavt7/nixedu-backend/badges/quality-score.png?b=master)
[![Code Coverage](https://scrutinizer-ci.com/g/asavt7/nixedu-backend/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/asavt7/nixedu-backend/?branch=master)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/asavt7/nixedu-backend)
![Codacy Grade](https://img.shields.io/codacy/grade/c9467ed47e064b1981e53862d0286d65)
![Github Repository Size](https://img.shields.io/github/repo-size/asavt7/nixedu-backend)
![Github Open Issues](https://img.shields.io/github/issues/asavt7/nixedu-backend)
![Lines of code](https://img.shields.io/tokei/lines/github/asavt7/nixedu-backend)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/asavt7/nixedu-backend)
![GitHub contributors](https://img.shields.io/github/contributors/asavt7/nixedu-backend)


# Решение задач курса [education.nixsolutions.com по Golang](https://education.nixsolutions.com/course/view.php?id=12)

## Требования

- ✅ Настроить golint в Goland IDE (используем https://github.com/mgechev/revive,
  т.к. https://github.com/golang/go/issues/38968)
- ✅ CRUD для постов и комментов. Ответ от АПИ должен реализовать 2 формата представления данных: XML, JSON.
- ❌ GORM
- ✅ Перепиши реализованное АПИ с использованием echo фреймворка
- ✅ Swagger specification
- ✅ JWT
- ✅ OAuth 2.0 (Google)
- ✅ Тестирование

## Дополнительные задания (от себя)
- ✅ Добавить docker-compose для запуска окружения для локальной разработки
- ✅ Dockerfile
- Добавить конфигурирование приложения (например, https://github.com/spf13/viper#readme)
- Фронт (React)
- Настроить CI/CD

## Использованные технологии

- Web framework https://github.com/labstack/echo
- Валидация https://github.com/go-playground/validator

- JWT https://jwt.io/
- Кеш для JWT-токенов https://redis.io/

- Миграции БД https://github.com/golang-migrate/migrate

- Документация https://github.com/swaggo/swag

- Mock-generator https://github.com/golang/mock
- Тестирование json ответа https://github.com/kinbiko/jsonassert

- Linter https://github.com/mgechev/revive


## Окружение для разработки

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