CREATE TABLE IF NOT EXISTS todos (
    id TEXT NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    dodate TEXT,
    finished INTEGER NOT NULL
);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("1", "Todo 1", "Description of this todo", "2025-08-25", 0);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("2", "Todo 2", "Description of this todo", "2025-08-25", 0);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("3", "Todo 3", "Description of this todo", "2025-08-25", 1);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("4", "Todo 4", "Description of this todo", "2025-08-25", 0);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("5", "Todo 5", "Description of this todo", "2025-08-25", 0);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("6", "Todo 6", "Description of this todo", "2025-08-25", 1);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("7", "Todo 7", "Description of this todo", "2025-08-25", 0);

INSERT INTO todos (id, title, description, dodate, finished)
VALUES ("8", "Todo 8", "Description of this todo", "2025-08-25", 0);
