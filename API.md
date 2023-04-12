# Application Programming Interface Documentation

## List
| No |     Web Service     | Method | URL |
|----|---------------------|--------|-----|
| 1 | [Create Product](#create-product) | POST | /api/products |
| 2 | [Delete Product](#delete-product) | DELETE | /api/products |
| 3 | [Get Product](#get-product) | GET | /api/products |


## Create Product
### URL : `/api/products`
### Method : `POST`

### Query Param
    request_id=101

### Header
    'X-API-Key: api-key-rahasia' \
    'Content-Type: application/json'

### Body Request
```json
{
    "customer" : "ryan",
    "price" : 12000,
    "quantity" : 20,
}
```

### Example cURL
curl --location 'localhost:3000/api/products?request_id=101' \
--header 'x-API-Key: api-key-rahasia' \
--header 'Content-Type: application/json' \
--data '{
    "customer" : "ryan",
    "price" : 12000,
    "quantity" : 20
}'

### Body Response Success
```json
{
    "code": 200,
    "status": "OK",
    "request_id": 101,
    "data": {
        "id": 20,
        "customer": "ryan",
        "price": 12000,
        "quantity": 20,
        "timestamp": "2023-04-12T00:23:06.396224+07:00"
    }
}
```

### Body Response Fail Unauthorized
```json
{
    "code": 401,
    "status": "UNAUTHORIZED",
    "request_id": 101,
    "data": null
}
```

## Delete Product
### URL : `/api/products`
### Method : `DELETE`

### Query Param
    request_id=101
    id=1

### Header
    'X-API-Key: api-key-rahasia' \
    'Content-Type: application/json'

### Example cURL
curl --location --request DELETE 'localhost:3000/api/products?request_id=101&id=1' \
--header 'x-API-Key: api-key-rahasia'

### Body Response Success
```json
{
    "code": 200,
    "status": "OK",
    "request_id": 101,
    "data": null
```

### Body Response Fail Unauthorized
```json
{
    "code": 401,
    "status": "UNAUTHORIZED",
    "request_id": 101,
    "data": null
}
```



