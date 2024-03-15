# curr-quote

## Описание

Данный проект представляет сервис для получения и обновления котировок валют.

Пользователь может получить имеющуюся наиболее свежую котировку, либо сделать 
запрос на обновление котировки, в ответ на него он получит id, с помощью 
которого через некоторое время сможет получить обновлённую котировку.

## Бизнес-логика

Swagger-документация к проекту доступна по адресу
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Обновление котировки

Пользователь в параметрах задаёт коды валют, котировку по которым он хотел бы 
обновить (например `EUR/MXN`), в теле ответа получает id, по которому можно 
получить обновлённую котировку

### Получение котировки по id

Пользователь в URL указывает id, полученный в ответ на запрос на обновление, в 
ответ получает отношение одной валюты к другой. Например, если был запрос на 
обновление `EUR/MXN`, то в ответе на запрос на получение обновлённой котировки 
будет `18.17815572`, то есть 1 евро оценивается примерно в 18 песо 
(на момент 16.03.2024).

### Получение последней доступной котировки

Пользователь в параметрах задаёт коды валют, котировку по которым он хотел бы
обновить (например `EUR/MXN`), в ответе получает отношение одной валюты к 
другой в том же формате, что и при получении котировки по id. Котировки 
автоматически обновляются раз в час, либо когда их обновляет сам пользователь.

## Детали реализации

Для получения котировок используется [exchange-api](https://github.com/fawazahmed0/exchange-api).

Котировки в приложении автоматически обновляются раз в час (*несмотря на это, 
информация всё равно будет устаревшей, так как котировки в API обновляются раз 
в сутки, зато бесплатно🙂*).

В качестве хранилища для котировок, для которых пользователи запрашивали 
обновление, используется PostgreSQL.

В качестве web-фреймворка использовался [gin](https://github.com/gin-gonic/gin).
Также для сервиса есть swagger-документация, доступная по адресу
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Бизнес-логика приложения покрыта unit-тестами, присутствует Docker-контейнеризация.

## Инструкция по запуску

### Запуск приложения

```shell
make run
```

Или

```shell
docker-compose up
```

### Запуск тестов

```shell
make test
```


