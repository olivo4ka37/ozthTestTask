# OzonTech Test Task

Этот репозиторий содержит исходный код для тестового задания OzonTech. Проект демонстрирует основные операции по работе с постами и комментариями с использованием базы данных PostgreSQL и GraphQL API.

Запустите приложение командой `make run`


### Примеры запросов
1. 
```
query {
  post(id: 1){
    id
    title
    content
    commentsEnabled
    author
  }
}
```

2.
```
query {
  posts{
    id
    title
    content
    commentsEnabled
    author
  }
}
```
3.
```
query {
	comment(postId: 1) {
    id
    author
    content
  }
}
```
4.
```
query {
	comments(postId: 1) {
    id
    author
    content
  }
}
```
5.
```
mutation {
  createPost(title: "is title", content:"is content",author:"I'm author") 
  {
    title
    content
    author
  }
}
```
6.
```
mutation {
  createComment(postId: 1, content: "НЕ ГРУСТИ",author: "РОСГОССТРАХ") {
    postId
    content
    author
  }
}
```
7.
```
mutation {
  updatePost(id: 1,title: "is title", content: "here is content"){
    id
    title
    content
  }
}
```
8.
```
mutation{
  updateComment(id: 1, content: "here is ANOTHER content") {
    id
    author
    content
  }
}
```