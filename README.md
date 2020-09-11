# ncanode-go

Клиент NCANode для Go

## Версия 1.0

Методы для работы с API v1.0. Смотреть документацию [тут](https://ncanode.kz/docs.php?go=b1b9b63034d9d4079aa08a214b4849255d28e6ab).

### Установка

```sh
go get -u github.com/danikarik/ncanode-go
```

### Использование

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

log.Println(resp.Result)
```

## Версия 2.0

Методы для работы с API v2.0. Смотреть документацию [тут](https://ncanode.kz/docs.php?go=3aed0a9b1d7f30ca5566f457fa0b2936f417fbd9).

### Установка

```sh
go get -u github.com/danikarik/ncanode-go/v2
```

### Использование

```go
import "github.com/danikarik/ncanode-go/v2"

client, err := ncanode.NewClient("http://127.0.0.1:14579")
if err != nil {
    log.Fatal(err)
}

resp, err := client.NodeInfo()
if err != nil {
    log.Fatal(err)
}

log.Println(resp.Result)
```

## Авторы

- [danikarik](https://github.com/danikarik)

## Лицензия

Проект лицензирован под [MIT](LICENSE)
