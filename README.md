# Cloud Run Weather Application

## Descrição
Esta aplicação fornece informações sobre o clima utilizando a API do OpenWeather. É uma aplicação escrita em Go e pode ser executada em um ambiente Docker.

## Formato do Request
Os requests devem ser feitos para o endpoint `/weather` com os seguintes parâmetros:
- `city`: Nome da cidade para a qual você deseja obter informações sobre o clima.
- `apikey`: Sua chave de API do OpenWeather.

### Exemplo de Request
```http
GET /weather?city=London&apikey=YOUR_API_KEY
```

## Buildando a Aplicação no Docker
Para construir a imagem Docker, execute o seguinte comando no diretório raiz do projeto:
```bash
docker build -t gabrielpgava/weather .
```

## Rodando a Aplicação no Docker
Após a construção da imagem, você pode rodar a aplicação com o seguinte comando:
```bash
docker run -p 8080:8080 gabrielpgava/weather:latest -d 
```

A aplicação estará disponível em `http://localhost:8080`.

