# MongoDB com Docker

Este projeto utiliza o MongoDB como banco de dados, executado em um container Docker.

## Como iniciar o MongoDB com Docker

1. Certifique-se de ter o Docker instalado em sua máquina.
2. Execute o comando abaixo para iniciar o serviço MongoDB:

```bash
docker-compose up -d
```

Isso irá iniciar o MongoDB em segundo plano, conforme definido no arquivo `docker-compose.yml`.

## Configuração padrão
- O serviço MongoDB será iniciado na porta padrão 27017.
- Os dados persistem no diretório `mongo-data/`.

## Parar o serviço
Para parar o MongoDB:

```bash
docker-compose down
```

## Observações
- Certifique-se de que a porta 27017 esteja livre.
- Os dados do banco são mantidos em `mongo-data/` mesmo após parar o container.

