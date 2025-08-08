# ---- Stage 1: Server Builder ----
FROM golang:1.22 AS serverbuilder

COPY ./server/src/go.mod ./

COPY ./server/src/go.sum ./

RUN ["go", "mod", "download"]

COPY ./server/src ./

RUN ["go", "build", "-o", "server", "."]


# ---- Stage 2: Client Builder ----
FROM node:24 AS clientbuilder

COPY ./client/package*.json ./

RUN npm install

COPY ./client ./

RUN ["npm", "run", "build"]


# ---- Stage 3: Run server ----
FROM ubuntu:22.04

COPY --from=serverbuilder . ./server
COPY --from=clientbuilder ./dist ./dist
COPY ./server/database/ToDoDatabase.db ./database

EXPOSE 8080

CMD ["./server/server"]