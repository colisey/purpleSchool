POST /v1/auth/login
POST /v1/auth/register
POST /v1/auth/restore

GET /v1/users/my // user from token
PUT /v1/users/my // user from token
GET /v1/users/my/posts // user from token
GET /v1/users/{id}/posts
GET /v1/users/my/comments // user from token
GET /v1/users/{id}/comments

GET /v1/posts
POST /v1/posts
PUT /v1/posts/{id}
DELETE /v1/posts/{id}
<!-- GET /v1/posts/byUser/{userId} -->

GET /v1/posts/{id}/comments
POST /v1/posts/{id}/comments
PUT /v1/posts/{id}/comments/{commentId}
DELETE /v1/posts/{id}/comments/{commentId}