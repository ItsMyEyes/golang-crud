# golang-crud

# Route
`Base Url: https://crud.kiyora-dev.xyz/` <br/>
`Authorization: Bearer <token>`
```
GET    /api/v1/home/          
GET    /api/v1/todo/          
POST   /api/v1/todo/          {title: "<title>", completed: <bool>} 
GET    /api/v1/todo/:id       
DELETE /api/v1/todo/:id       
POST   /api/v1/user/register  {username: "<string>", email: "<string>", password: "<string>"}
POST   /api/v1/user/login     {username: "<string>", password: <string>} 
GET    /api/v1/user/logout    
GET    /api/v1/               
```

# Installation
1. Clone this repository
```
  $ git clone https://github.com/ItsMyEyes/golang-crud
 ```
2. Change directory to this repository
```
  $ cd golang-crud
 ```

3. Install all dependencies
```
 $ go mod download
 ```

4. set environment variable
```
 $ cp .env.example .env
 ```

5. Run the server
```
 $ go run main.go
 ```

# Usage

## Register
```
POST /api/v1/user/register
```

Request Body
```
{
  "username": "<string>",
  "email": "<string>",
  "password": "<string>"
}
```

Response
```
{
  "status": "success",
  "message": "Register success",
  "data": {
    "id": "<int>",
    "username": "<string>",
    "email": "<string>",
    "password": "<string>",
    "created_at": "<time>",
    "updated_at": "<time>"
  }
}
```
