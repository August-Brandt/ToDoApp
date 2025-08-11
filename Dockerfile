FROM ubuntu:22.04

WORKDIR /app

COPY ./server/src/server ./server/src/
COPY ./client/dist ./client/dist/
COPY ./server/database/ToDoDatabase.db ./server/database/

EXPOSE 8080

WORKDIR /app/server/src

CMD ["./server"]