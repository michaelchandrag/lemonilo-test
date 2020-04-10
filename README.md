# lemonilo-test

### Installation
clone this repo under `{GOPATH}/src/github.com/michaelchandrag/lemonilo-test`

set up MySQL database, dump table can be found under `database/lemonilo.sql`

set up environment variable at `.env`


run `./run`

### Environment Variable
```
PORT=8080
DB_HOST=localhost
DB_USER=root
DB_PASS=
DB_NAME=lemonilo
```

### Database
database structure can be found under
`database/lemonilo.sql`


### Endpoint
Read
```
URL: [GET] {BASE_URL}/user_profiles
Body Response:
{
    "data": [
        {
            "user_id": 1,
            "email": "canzinzzzide@yahoo.co.id",
            "password": "--",
            "address": "address"
        },
        {
            "user_id": 3,
            "email": "michaelchandrag114@yahoo.co.id",
            "password": "--",
            "address": "address"
        }
    ],
    "success": true
}
```

Create
```
URL: [POST] {BASE_URL}/user_profile
Body Request:
{
	"email": "sample@baru.com",
	"password": "baru",
  "address": "address"
}
Body Response:
{
    "data": {
        "user_id": 8,
        "email": "baru@baru.com",
        "password": "5ef035d11d74713fcb36f2df26aa7c3d",
        "address": ""
    },
    "message": "Success create user profile.",
    "success": true
}
```

Update
```
URL: [PUT] {BASE_URL}/user_profile/:user_id
Body Request:
{
	"email": "sample@baru.com",
	"password": "baru",
  "address": "address"
}
Body Response:
{
    "message": "Update user profile success.",
    "success": true
}
```

Delete
```
URL: [DELETE] {BASE_URL}/user_profile/:user_id
Body Response:
{
    "message": "Delete user profile success",
    "success": true
}
```

Login
```
URL: [POST] {BASE_URL}/login
Body Request:
{
	"email": "baru@baru.com",
	"password": "baru"
}
Body Response:
{
    "data": {
        "user_id": 9,
        "email": "baru@baru.com",
        "password": "5ef035d11d74713fcb36f2df26aa7c3d",
        "address": "asd"
    },
    "message": "Success login.",
    "success": true
}
```
