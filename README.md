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
POST   /api/v1/user/register  {name: "<string>", email: "<string>", password: "<string>"}
POST   /api/v1/user/login     {email: "<title>", password: <bool>} 
GET    /api/v1/user/logout    
GET    /api/v1/               
```

```
  git clone https://github.com/ItsMyEyes/golang-crud //repo git
  code golang-crud // open vsc studio
  go mod tidy //for installation package
  copy .env.example .env 
  go run main.go //to run dev
  go buil . //to build
 ```
