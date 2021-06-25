# rest-api-with-unit-test
It contains Rest api with basic unit test.


**steps to use:**
1. Clone the project using `git clone --recursive https://github.com/manoj-JU/rest-api-with-unit-test.git`
2. get inside the repo : `cd rest-api-with-unit-test`
3. start the service using : `go run main.go`
4. run unit test using `go test -v`
5. call APIs.
  
# APIs
1.**Add user**

**URL:** `http://localhost:4563/api/v1/users`

**Method:** POST

**Content-Type:** application/json

**Authentication Required:** NO

**Data Examples:**

**Query Params:**

```json
{
"user_name": "Manoj",
"phone_number": "23123123123"
}
```

## Success Response:

**Code:** `200 Ok`


**Content Example:**

```json
{
    "user_id": "902081",
    "user_name": "Manoj",
    "phone_number": "23123123123"
}
```

2.**Get all users**

**URL:** `http://localhost:4563/api/v1/users`

**Method:** GET

**Content-Type:** application/json

**Authentication Required:** NO

**Data Examples:**

## Success Response:

**Code:** `200 OK`


**Content Example:**

```json
[
    {
        "user_id": "498081",
        "user_name": "rohit",
        "phone_number": "2232232323"
    },
    {
        "user_id": "727887",
        "user_name": "manoj",
        "phone_number": "312312312"
    }
]
```

3.**Get a user using user id**

**URL:** `http://localhost:4563/api/v1/users/{id}`

**Method:** GET

**Content-Type:** application/json

**Authentication Required:** NO

## Success Response:

**Code:** `200 Ok`


**Content Example:**

```json
 {
    "user_id": "727887",
    "user_name": "manoj",
    "phone_number": "312312312"
}
```

## Error Response:

**Condition:** If Invalid user id provided.

**Code:** `400 Bad Request`

**Content Example:**

```json:
{
    "message": "Invalid id"
}
```

