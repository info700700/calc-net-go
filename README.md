# calc-net-go
Калькулятор  
Клиент-серверное приложение

--- 

Запуск приложения в виде исполняемого файла.
```
$ make
$ ./calc-server
```

Запуск приложения в докере.
```
$ make build_image
$ docker run --name calc-server --detach --publish 80:80 calc-server
```