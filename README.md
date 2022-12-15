# rest-api-tutorial

# user-service

# REST API

GET /users -- list of users -- 200, 404, 500
GET /users:id -- user by id -- 200, 404, 500
POST /users:id -- create user -- 201(HTTP 201 Created Код ответа об успешном статусе указывает, что запрос выполнен успешно и привёл к созданию ресурса ), 4xx , Header Location: url
PUT /users:id -- full update users --204/200 (в body полностью обновленный вользователь)
DELETE /users:if --delete update user --204/200, 404, 500
PATCH /USERS:id --partially update user --204/200 (либо ничего не отдаем( No Content ), либо в body полностью обновленный вользователь) 