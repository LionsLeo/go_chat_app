# go_chat_app
This is a full stack project for a room based chat application the tech stack of this project is Go lang, Postgres and Flutter.

## Commands to execute
docker pull postgres:15-alpine
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine
docker exec -it postgres15 createdb --username=root --owner=root go-chat
docker exec -it postgres15 psql