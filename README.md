# Golang Banking API

## Endpoints

| Endpoint | READ | CREATE | EDIT | DELETE |
| :---     |:----:|:------:|:----:|:------:|
| /login | ❌ | ✅ | ❌ | ❌ |
| /account | ✅ | ✅ | ❌ | ❌ |
| /account /**{id}** | ✅ | ❌ | ❌ | ✅ |
| /account /**{id}** /transfer | ❌ | ✅ | ✅ | ❌ |
