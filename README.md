# Просто разбор интересных примеров

## Темы

### Слайсы

### Каналы

### Мапы

### For

### Interfaces

### Goroutines

### Figures(Draw Gif)(Донован)

### Fetch (Донован 1.8, 1.9)

Запуск в консоле через

```go
go run fetch_donov.go http://www.google.com
```

### Fetch ALL (Донован)

Запуск в консоле через

```go
go run fetchAll_donovan.go https://golang.org http://gopl.io https://godoc.org
```

### echo-server-with-mutex (localhost:8000)

```go
 go run ./echo-server-with-mutex.go
```

### flagsCommandLine

```go
  go run ./flagsCommandLine.go / a bc def
```

### tempConv - Конвертация температуры

```go
  go run ./tempConv.go
```

### Cf - Конвертация температуры

```go
  go run ./CfDonovan.go
```

## lesson 4 - WEB VK-TEAM

## lesson 4 - Telegram Bot VK-TEAM

### BitOperators - побитовые операции

```go
  go run ./main.go
```

### Benchmarks

```go
go test -bench=.
```

### Остановка Контекста

#### Остановка канала

```go
go run ./contextEx/contextEx.go
```

#### Errors - создание кастомных ошибок

```go
go run ./errors/errors.go
```

### включить поддержку модулей в нашей новой программе.

```go
go mod init mod

go mod tidy
```

#### HTMLServer11Lessons

Простой Html Сервер

```go
go run htmlServer11projects/main.go
```

### visitLinks

```go
go run fetch_donov.go https://golang.org | go run ./visitLinksDonovan/main.go
```

### outline

```go
go run fetch_donov.go https://golang.org | go run ./outlineDonovan/main.go
```

### outline2

#### Напишите функцию для заполнения отображения, ключами которого являются имена элементов (р, d iv , sp an и т.д.), а значениями — количество элементов с таким именем в дереве HTML-документа

```go
go run fetch_donov.go https://golang.org | go run ./outline2Donovan/main.go
```

#### Напишите функцию для вывода содержимого всех текстовых узлов в дереве документа HTML. Не входите в элементы <script> и <style> , поскольку их содержимое в веб-браузере не является вид

```go
go run fetch_donov.go https://golang.org | go run ./outline3Donovan/main.go
```


####  crudServer11projects

####  MysqlServer11projects

####  SlackBot11projects
##### Пока не запускал - нужен API

####  SlackBotFileUploading11projects
##### Пока не запускал - нужен API

### FlagCli (70 techniques)
go run flagCli/main.go
go run flagCli/main.go -s -name buttercup

