🌚🌚🌚🌚🌚🌚
=============
Что умеет
----------
1. Слать Пашка (/pashok)
2. Отвечать на вопросы (/ask)

Как установить
--------------
```bash
go get bitbucket.org/MegaDeathLightsaber/bubnobot
```

Потом в исходном коде вставить свой токен и привязку к чату:
```go
const (
	// Token to access HTTP Bot API
	apiToken  = "1337:dyadyabogdan"
	ourChatID = -111222333444555
)
```

А также URL прокси-сервера:
```go
const proxyStr = "yourproxy"
```