# ncanode-go

Клиент NCANode для Go

## Установка

```sh
go get -u github.com/danikarik/ncanode-go
```

## Использование

```go
import "github.com/danikarik/ncanode-go"

client, err := ncanode.NewClient("http://127.0.0.1:14579")
if err != nil {
    log.Fatal(err)
}

resp, err := client.NodeInfo()
if err != nil {
    log.Fatal(err)
}

log.Println(resp)
```

## Авторы

- [danikarik](https://github.com/danikarik)

## Лицензия

Проект лицензирован под [MIT](LICENSE)
