FROM node:lts-alpine AS frontend
WORKDIR /app

COPY ./client/package.json ./client/package.json

WORKDIR /app/client

RUN ["npm", "i"]

COPY client /app/client/

RUN ["npm", "run", "build"]


FROM golang:1.22 AS server
WORKDIR /app

COPY ./server/src ./server

WORKDIR /app/server

RUN ["go", "mod", "tidy"]

RUN ["go", "build", "-o", "server"]


FROM ubuntu:latest
WORKDIR /TodoApp/app

COPY --from=frontend /app/client/dist ./client/dist
COPY --from=server /app/server/server ./server/src/
COPY ./server/database/ToDoDatabase.db ./server/database/

WORKDIR /TodoApp/app/server/src

EXPOSE 8080

CMD ["./server"]