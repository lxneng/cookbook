# hello-gin-jwt

## apis

```
❯ echo '{"Email": "hey@lxneng.com", "Password": "serect", "Name": "Eric"}' | http POST localhost:8100/api/v1/signup
HTTP/1.1 200 OK
Content-Length: 238
Content-Type: application/json; charset=utf-8
Date: Tue, 22 Dec 2020 07:05:13 GMT

{
    "CreatedAt": "2020-12-22T15:05:13.643136965+08:00",
    "DeletedAt": null,
    "ID": 1,
    "UpdatedAt": "2020-12-22T15:05:13.643136965+08:00",
    "email": "hey@lxneng.com",
    "name": "Eric",
    "password": "$2a$14$uu3hYcLyjKgGmxx7CtXCIOt4rBq6klFqDIssqFZKhQutovL/pLeiy"
}


❯ echo '{"Email": "hey@lxneng.com", "Password": "serect"}' | http POST localhost:8100/api/v1/signin
HTTP/1.1 200 OK
Content-Length: 177
Content-Type: application/json; charset=utf-8
Date: Tue, 22 Dec 2020 07:06:56 GMT

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImhleUBseG5lbmcuY29tIiwiZXhwIjoxNjA4NzA3MjE2LCJpc3MiOiJBdXRoU2VydmljZSJ9.5rR0l5JYpwmIVLW98QTQQU1rdGZ59klPTmDNIBHhYvU"
}

❯ http localhost:8100/api/v1/profile Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImhleUBseG5lbmcuY29tIiwiZXhwIjoxNjA4NzA3MjE2LCJpc3MiOiJBdXRoU2VydmljZSJ9.5rR0l5JYpwmIVLW98QTQQU1rdGZ59klPTmDNIBHhYvU"
HTTP/1.1 200 OK
Content-Length: 178
Content-Type: application/json; charset=utf-8
Date: Tue, 22 Dec 2020 07:13:12 GMT

{
    "CreatedAt": "2020-12-22T15:05:13.643136965+08:00",
    "DeletedAt": null,
    "ID": 1,
    "UpdatedAt": "2020-12-22T15:05:13.643136965+08:00",
    "email": "hey@lxneng.com",
    "name": "Eric",
    "password": ""
}

```