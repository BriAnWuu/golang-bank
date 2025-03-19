# Golang Banking API

## Endpoints

| Endpoint | READ | CREATE | EDIT | DELETE | AUTH |
| :---     |:----:|:------:|:----:|:------:|:----:|
| /login | ❌ | ✅ | ❌ | ❌ | |
| /account | ✅ | ✅ | ❌ | ❌ | |
| /account /**{id}** | ✅ | ❌ | ❌ | ✅ | ✅ |
| /account /**{id}** /transfer | ❌ | ✅ | ✅ | ❌ | ✅ |


## Manage Postgres database using Docker
Create and run your Postgres image with Docker
```
$ docker run --name <CONTAINER_NAME> -e POSTGRES_PASSWORD=<ENV_PASSWORD> -p 5432:5432 -d postgres

```

Use `$ Docker ps` to show running containers