# Desafio técnico Go Expert FullCycle

## Execução

`go run main.go  --url=http://google.com --requests=25 --concurrency=4`

## Executando com o Docker

`docker run --rm -it $(docker build -q .) --url=http://google.com --requests=25 --concurrency=4`
Isso fará o build da imagem, execução e quando concluído será deletada automaticamente.