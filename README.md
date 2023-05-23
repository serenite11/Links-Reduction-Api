# Links-Reduction-Api
 
Для запуска сервиса клонируйте репозиторий и выполните команду:
 
`make compose`


Для выбора хранилища в файле docker-compose.yml мжно изменить переменную STORE :

`STORE="POSTGRES"/"IN-MEMORY"`

## REST

### POST

`http://localhost:8080/`

<img src="imgs/Post-Rest.jpg" width="" height="">


### GET

`http://localhost:8080/:link` 

<img src="imgs/Get-Rest.jpg" width="" height="">


## gRPC

`localhost:5500`

### CreateShortUrl

<img src="imgs/CreateShortUrl.jpg" width="" height="">

### GetLongUrl

<img src="imgs/GetLongUrl.jpg" width="" height="">

Для тестов используйте команду: \
`make test`