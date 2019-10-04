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

Как запустить
-------------
1. Заходим в папку с исходным кодом, в каталог cmd/bubnobot.
2. Пишем:
	```bash
	go install
	```
3. Ура, можем запускать при из консоли, набирая
   ```bash
   bubnobot
   ```