# parking-project

## :rocket: Run App
### executando aplicação:
```sh
go run main.go
```

## :books: Endpoints API
* ### `GET` `localhost:8080/parking-info`
* _Descrição_: Endpoint para listar quantidade livres e ocupadas do estacionamento, e quantidades de veículos estacionados
    * `RESPONSE`
  ```JSON
    {
      "free_parking_spaces": 100,
      "occupied_parking_spaces": 0,
      "vehicle_information": {
        "car": 0,
        "motorbike": 0,
        "van": 0
      }
    }
    ```
  
* ### `POST` `localhost:8080/parking/occupy`
* _Descrição_: Endpoint para ocupar a vaga livre do estacionamento
* _:vehicle_: Opções: **car / motorbike / van**
    * `Body`
  ```JSON
    {
      "name": ""
    }
    ```
* ### `POST` `localhost:8080/parking/release`
* _Descrição_: Endpoint para desacupar uma vaga e voltar a ficar livre
* _:vehicle_: Opções: **car / motorbike / van**
    * `Body`
  ```JSON
    {
      "name": ""
    }
    ```