# graphql-example

## login
```shell
curl -X POST "http://127.0.0.1:8000/login" -H "Content-Type: application/json" -d '{"name":"user1","password":"password1"}' -v
```

```
...
< HTTP/1.1 200 OK
< Set-Cookie: session=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk4OTg1MzUsInN1YiI6IjA0YjRkMzQzLWQ5NzQtNDUxZS1iYjE0LWNjNTc1NTIwNGM3YSJ9.DrkTVdrfy5p0iBNW_aHFUFQLyxtq58Xd-eYYxH_3LDs
...
```

```shell
export session=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk4OTg1MzUsInN1YiI6IjA0YjRkMzQzLWQ5NzQtNDUxZS1iYjE0LWNjNTc1NTIwNGM3YSJ9.DrkTVdrfy5p0iBNW_aHFUFQLyxtq58Xd-eYYxH_3LDs
```

## getPost

```graphql
query{
    getPost(id:"b711966f-1a9d-46d5-8f5e-c2bfe8f94229"){
        id,
        title,
        content,
        user{
            id,
            name,
            posts{
                id,
                title,
                content
            }
        }
    }
}
```

```shell
curl -s -X POST "http://127.0.0.1:8000/graphql" -H "Content-Type: application/json" -H "Cookie: session=$session" -d '{"query":"query{getPost(id:\"b711966f-1a9d-46d5-8f5e-c2bfe8f94229\"){id,title,content,user{id,name,posts{id,title,content}}}}"}' | jq
```
```json
{
  "data": {
    "getPost": {
      "content": "content1",
      "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
      "title": "title1",
      "user": {
        "id": "04b4d343-d974-451e-bb14-cc5755204c7a",
        "name": "user1",
        "posts": [
          {
            "content": "content1",
            "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
            "title": "title1"
          },
          {
            "content": "content2",
            "id": "0ae944de-40a9-447d-9066-89d3a190e32f",
            "title": "title2"
          }
        ]
      }
    }
  }
}
```

## getPosts

```graphql
query{
    getPosts{
        id,
        title,
        content,
        user{
            id,
            name,
            posts{
                id,
                title,
                content
            }
        }
    }
}
```

```shell
curl -s -X POST "http://127.0.0.1:8000/graphql" -H "Content-Type: application/json" -H "Cookie: session=$session" -d '{"query":"query{getPosts{id,title,content,user{id,name,posts{id,title,content}}}}"}' | jq
```

```json
{
  "data": {
    "getPosts": [
      {
        "content": "content1",
        "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
        "title": "title1",
        "user": {
          "id": "04b4d343-d974-451e-bb14-cc5755204c7a",
          "name": "user1",
          "posts": [
            {
              "content": "content1",
              "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
              "title": "title1"
            },
            {
              "content": "content2",
              "id": "0ae944de-40a9-447d-9066-89d3a190e32f",
              "title": "title2"
            }
          ]
        }
      },
      {
        "content": "content2",
        "id": "0ae944de-40a9-447d-9066-89d3a190e32f",
        "title": "title2",
        "user": {
          "id": "04b4d343-d974-451e-bb14-cc5755204c7a",
          "name": "user1",
          "posts": [
            {
              "content": "content1",
              "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
              "title": "title1"
            },
            {
              "content": "content2",
              "id": "0ae944de-40a9-447d-9066-89d3a190e32f",
              "title": "title2"
            }
          ]
        }
      },
      {
        "content": "content3",
        "id": "ac95a094-82ba-47ed-ae84-f46053816544",
        "title": "title3",
        "user": {
          "id": "67ed72e2-7ffb-4427-9450-b87aab994703",
          "name": "user2",
          "posts": [
            {
              "content": "content3",
              "id": "ac95a094-82ba-47ed-ae84-f46053816544",
              "title": "title3"
            },
            {
              "content": "content4",
              "id": "526875df-18f3-44ae-b60e-908a935e605f",
              "title": "title4"
            }
          ]
        }
      },
      {
        "content": "content4",
        "id": "526875df-18f3-44ae-b60e-908a935e605f",
        "title": "title4",
        "user": {
          "id": "67ed72e2-7ffb-4427-9450-b87aab994703",
          "name": "user2",
          "posts": [
            {
              "content": "content3",
              "id": "ac95a094-82ba-47ed-ae84-f46053816544",
              "title": "title3"
            },
            {
              "content": "content4",
              "id": "526875df-18f3-44ae-b60e-908a935e605f",
              "title": "title4"
            }
          ]
        }
      }
    ]
  }
}
```

## createPost

```graphql
mutation{
    createPost(title:"title5",content:"content5"){
        id,
        title,
        content,
        user{
            id,
            name,
            posts{
                id,
                title,
                content
            }
        }
    }
}
```

```shell
curl -s -X POST "http://127.0.0.1:8000/graphql" -H "Content-Type: application/json" -H "Cookie: session=$session" -d '{"query":"mutation{createPost(title:\"title5\",content:\"content5\"){id,title,content,user{id,name,posts{id,title,content}}}}"}' | jq
```

```json
{
  "data": {
    "createPost": {
      "content": "content5",
      "id": "4ae67d5f-c92d-477a-b7b5-90412bb0a65c",
      "title": "title5",
      "user": {
        "id": "04b4d343-d974-451e-bb14-cc5755204c7a",
        "name": "user1",
        "posts": [
          {
            "content": "content1",
            "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
            "title": "title1"
          },
          {
            "content": "content2",
            "id": "0ae944de-40a9-447d-9066-89d3a190e32f",
            "title": "title2"
          },
          {
            "content": "content5",
            "id": "4ae67d5f-c92d-477a-b7b5-90412bb0a65c",
            "title": "title5"
          }
        ]
      }
    }
  }
}
```

## updatePost

```graphql
mutation{
    updatePost(id:"4ae67d5f-c92d-477a-b7b5-90412bb0a65c",title:"updated-title5",content:"updated-content5"){
        id,
        title,
        content,
        user{
            id,
            name,
            posts{
                id,
                title,
                content
            }
        }
    }
}
```

```shell
curl -s -X POST "http://127.0.0.1:8000/graphql" -H "Content-Type: application/json" -H "Cookie: session=$session" -d '{"query":"mutation{updatePost(id:\"4ae67d5f-c92d-477a-b7b5-90412bb0a65c\",title:\"updated-title5\",content:\"updated-content5\"){id,title,content,user{id,name,posts{id,title,content}}}}"}' | jq
```

```json
{
  "data": {
    "updatePost": {
      "content": "updated-content5",
      "id": "4ae67d5f-c92d-477a-b7b5-90412bb0a65c",
      "title": "updated-title5",
      "user": {
        "id": "04b4d343-d974-451e-bb14-cc5755204c7a",
        "name": "user1",
        "posts": [
          {
            "content": "content1",
            "id": "b711966f-1a9d-46d5-8f5e-c2bfe8f94229",
            "title": "title1"
          },
          {
            "content": "content2",
            "id": "0ae944de-40a9-447d-9066-89d3a190e32f",
            "title": "title2"
          },
          {
            "content": "updated-content5",
            "id": "4ae67d5f-c92d-477a-b7b5-90412bb0a65c",
            "title": "updated-title5"
          }
        ]
      }
    }
  }
}
```

## deletePost

```graphql
mutation{
    deletePost(id:"4ae67d5f-c92d-477a-b7b5-90412bb0a65c")
}
```

```shell
curl -s -X POST "http://127.0.0.1:8000/graphql" -H "Content-Type: application/json" -H "Cookie: session=$session" -d '{"query":"mutation{deletePost(id:\"4ae67d5f-c92d-477a-b7b5-90412bb0a65c\")}"}' | jq
```

```json
{
  "data": {
    "deletePost": "4ae67d5f-c92d-477a-b7b5-90412bb0a65c"
  }
}
```
