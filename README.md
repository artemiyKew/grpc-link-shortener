
[![forthebadge](https://forthebadge.com/images/featured/featured-built-with-love.svg)](https://forthebadge.com)
# GRPC-gateway link shortener
Сокращатель ссылок

Тестовое задание 

Используемые технологии: 
- gRPC-gateway
- Docker (для запуска сервиса)

# Usage

**Скопируйте проект**
```bash
git clone https://github.com/artemiyKew/grpc-link-shortener.git
```

**Перейдите в каталог проекта**
```bash
cd grpc-link-shortener
```

**Запустите сервер**
```bash
make compose
```

## Examples
- [Создание сокращенной ссылки](#создание-сокращенной-ссылки)
- [Редирект](#редирект) 
- [Получение данных о сокращенных ссылках](#получение-данных-о-сокращенных-ссылках)

## Создание сокращенной ссылки
Создание сокращенной ссылки: 

```bash
curl -X POST \
    http://0.0.0.0:2222/ \
    -H "Content-Type: application/json" \
    -d '{
    "full_url": "https://youtube.com"
    }'
```
Пример ответа: 
```json
{
    "fullUrl": "https://youtube.com",
    "shortUrl": "4dc3a76939",
    "visitsCount": 0
}
```

## Редирект
Редирект:
```bash
curl http://0.0.0.0:2222/4dc3a76939 
```
К сожалению я не знаю, как правильно сделать redirect D;

Пример ответа: 
```json
{
    "fullUrl": "https://youtube.com",
}
```
## Получение данных о сокращенных ссылках
Получение данных о сокращенных ссылках:

```bash
curl -X GET \
    http://0.0.0.0:2222/ \
    -H "Content-Type: application/json"
```
Пример ответа: 
```json
{
    "result": [
        {
            "fullUrl": "https://vk.com",
            "shortUrl": "877ca9b8dd",
            "visitsCount": 1
        },
        {
            "fullUrl": "https://example.com",
            "shortUrl": "a379a6f6ee",
            "visitsCount": 0
        },
        {
            "fullUrl": "https://youtube.com",
            "shortUrl": "4dc3a76939",
            "visitsCount": 1
        }
    ] 
}
```



