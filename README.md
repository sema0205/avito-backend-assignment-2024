# Banner management service

Микросервис для управления баннерами, предназначен для создания, редактирования, удаления и получения баннеров с
возможностью фильтрации по различным параметрам. Администраторы могут управлять баннерами, включая их размещение и
обновление.

Используемые технологии:

- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- Gin (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- pgx (драйвер для работы с PostgreSQL)
- golang/mock, testify (для тестирования)

Сервис разработан с использованием принципов Clean Architecture, обеспечивая удобство в расширении функциональности и
упрощение тестирования. Дополнительно внедрён механизм Graceful Shutdown для безопасного завершения процессов сервиса.

# Getting Started

Для настройки приложения необходимо выполнить следующие шаги:

- Добавить файл `.env` в директорию с проектом, заполнив его данными на основе `.env.example`.
- По желанию, можно настроить параметры в файле `config/config.yaml` для конфигурации различных аспектов приложения.

Эти действия обеспечат корректную настройку приложения.

# Usage

- Запуск сервиса осуществляется командой `make compose-up`.

## Examples

Некоторые примеры запросов

- [Аутентификация](#sign-in)
- [Получение баннера для юзера](#user-banner)
- [Получение фильтрованных баннеров для админа](#admin-get-banners)
- [Создание баннера](#create-banner)
- [Обновление баннера](#update-banner)
- [Удаление баннера](#delete-banner)

### Аутентификация <a name="sign-in"></a>

Аутентификация сервиса в режиме юзера:

```curl
curl --location --request POST 'localhost:8080/api/v1/users/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "username":"test",
    "password":"test"
}'
```

Аутентификация сервиса в режиме админа:

```curl
curl --location --request POST 'localhost:8080/api/v1/admins/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "username":"test",
    "password":"test"
}'
```

Пример ответа:

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImlkIjoxLCJleHAiOjE3MTMxMDU4OTQsImlhdCI6MTcxMzA5ODY5NH0.mDfa31Fv1ahVYsVJtA9amVOAQ6fa213bc9bLWz4NpWA"
}
```

### Получение баннера для юзера <a name="user-banner"></a>

Получение баннера для юзера:

```curl
curl --location --request GET 'localhost:8080/api/v1/users/banner?tag_id=1&feature_id=1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJpZCI6MSwiZXhwIjoxNzEzMTA1OTc5LCJpYXQiOjE3MTMwOTg3Nzl9.8cGudN75lpdRUMbazmmN35xX-uBO0PuyH2tTtrCvO00'
```

Пример ответа:

```json
{
  "content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}"
}
```

### Получение фильтрованных баннеров для админа <a name="admin-get-banners"></a>

Получение фильтрованных баннеров для админа:

```curl
curl --location --request GET 'localhost:8080/api/v1/admins/banners?feature_id=1&tag_id=1&limit=10' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJpZCI6MSwiZXhwIjoxNzEzMTA1OTc5LCJpYXQiOjE3MTMwOTg3Nzl9.8cGudN75lpdRUMbazmmN35xX-uBO0PuyH2tTtrCvO00'
```

Пример ответа:

```json
[
  {
    "banner_id": 1,
    "content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}",
    "is_active": true,
    "updated_at": "2024-04-14T15:38:48.782437+03:00",
    "created_at": "2024-04-14T15:38:48.782437+03:00",
    "tag_ids": [
      1
    ],
    "feature_id": 1
  }
]
```

### Создание баннера <a name="create-banner"></a>

Создания баннера:

```curl
curl --location --request POST 'localhost:8080/api/v1/admins/banners' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJpZCI6MSwiZXhwIjoxNzEzMTA1OTc5LCJpYXQiOjE3MTMwOTg3Nzl9.8cGudN75lpdRUMbazmmN35xX-uBO0PuyH2tTtrCvO00' \
--header 'Content-Type: application/json' \
--data '{
  "tag_ids": [
    1
  ],
  "feature_id": 1,
  "content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}",
  "is_active": true
}'
```

Пример ответа:

```json
{
  "banner_id": 2
}
```

### Обновление баннера <a name="update-banner"></a>

Обновление баннера:

```curl
curl --location --request PATCH 'localhost:8080/api/v1/admins/banners/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJpZCI6MSwiZXhwIjoxNzEzMTA1OTc5LCJpYXQiOjE3MTMwOTg3Nzl9.8cGudN75lpdRUMbazmmN35xX-uBO0PuyH2tTtrCvO00' \
--header 'Content-Type: application/json' \
--data '{
  "tag_ids": [
    1
  ],
  "feature_id": 1,
  "content": "{\"new_title\": \"new_title\", \"text\": \"some_text\", \"url\": \"some_url\"}",
  "is_active": false
}'
```

### Удаление баннера <a name="delete-banner"></a>

Удаление баннера:

```curl
curl --location --request DELETE 'localhost:8080/api/v1/admins/banners/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJpZCI6MSwiZXhwIjoxNzEzMTA1OTc5LCJpYXQiOjE3MTMwOTg3Nzl9.8cGudN75lpdRUMbazmmN35xX-uBO0PuyH2tTtrCvO00'
```

# Decisions <a name="decisions"></a>

В процессе разработки пришлось принять несколько ключевых решений для оптимизации работы сервиса и улучшения его
функциональности:

1. **Разделение токенов для пользователей и администраторов:**

> Для улучшения безопасности и разделения доступа между разными типами пользователей было решено использовать
> поле `Role` в token claims, которое может принимать значения `admin` и `user`. Это позволяет легко разграничивать права
> доступа внутри сервиса и упрощает управление пользователями, обеспечивая администраторам более широкие возможности по
> сравнению с обычными пользователями.

2. **Интеграция in-memory кэша в сервис баннеров:**

> Для повышения производительности при получении данных о баннерах было решено внедрить in-memory кэш. Это позволило
> значительно ускорить процесс чтения информации, особенно при высокой нагрузке на сервис. Кэш реализован таким образом,
> что он автоматически обновляется при изменениях данных о баннерах, что обеспечивает актуальность информации без
> значительных задержек.

Эти решения были направлены на оптимизацию работы сервиса, повышение его стабильности и удобства использования как для
пользователей, так и для администраторов.