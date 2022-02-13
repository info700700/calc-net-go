# calc-net-go
Калькулятор  
Клиент-серверное приложение

--- 

Запуск приложения в докере.
```
$ make build_image
$ docker run --detach --publish 80:80 calc-server
```