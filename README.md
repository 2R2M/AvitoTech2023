# AvitoTech2023

 Сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

Используемые технологии:

- PostgreSQL 
- Gin
- SQLX 
- Docker 
- Swagger
- Slog
 

# Запуск

docker-compose build && docker-compose up

# Примеры запросов

### Добавление пользователя
```
curl --location 'localhost:8085/stat/v1/user' \
--header 'Content-Type: application/json' \
--data '{
    "id":"1"
}'
```

### Добавление сегмента

```
curl --location 'localhost:8085/stat/v1/segment' \
--header 'Content-Type: application/json' \
--data '{
    "slug":"AVITO_PERFORMANCE_TEST"
}'
```

### Удаление сегмента

```
curl --location --request DELETE 'localhost:8085/stat/v1/delete/AVITO_PERFORMANCE_TEST'
```

### Добавление пользователя в сегменты
```
curl --location 'localhost:8085/api/v1/users/1/segments' \
--header 'Content-Type: application/json' \
--data '{
    "slugs": [
        "AVITO_PERFORMANCE_TEST", "AVITO_PERFORMANCE_VAS"],
        "expired_at": "2006-01-02 15:04:05"
}'
```

### Удаление пользователей из сегментов
```
curl --location --request DELETE 'localhost:8085/stat/v1/users/1/segments' \
--header 'Content-Type: application/json' \
--data '{
    "slugs": [
       "AVITO_PERFORMANCE_TEST"]
}
'
```

### Получение сегментов пользователя
```
curl -X 'GET' \
  'http://localhost:8080/api/v1/users/1/segments' \
  -H 'accept: application/json'
```
### Получение ссылки на CSV файл
```
curl --location 'localhost:8085/api/v1/report/8/2023'
```

### Скачивание CSV файла
```
curl --location 'localhost:8085/download/report_2023_8.csv'
```
# Результат

В ходе выполнения задания были реализованы функционал основного задания, 1 и 2 дополнительное задание. Также были написаны тесты и создан swagger. Метод добавления и удаления пользователя из сегмента были разделены на 2 отдельных метода. 
Принцип работы: Существует функционал позволяющий создавать пользователей, создавать/удалять пользователей. Пользователя можно добавлять в сегменты и удаляться из сегментов. Для уменьшение ошибок была создана таблица Operation. Каждая операция (добавление и удаление) заносится в таблицу Operation. После асинхронная функиция каждую секунду проверять таблицу operation и создает запись в таблице UserSegment, либо удаляет ее. В случае удачного добавления/удаления каждая операция помечается как Success или Error. Для корректной работы, прежде чем "перекладывать" операции из Operation в  UserSegment, операции сортируются. Также для реализации 1 доп. задания используется таблица Operation из данных которой создаются отчеты (используются только данные, помечанные как SUCCESS). Для выполнения второго задания в UserSegment добавлено поле expired_at, которое мы проверяем, прежде чем сегменты конкретного пользователя. При удалении сегмента из базы данных мы создаем операции удаления для каждого пользователя, который входит в этот сегмент и только после этого удаляем сам сегмент из базы данных.
