DROP TABLE IF EXISTS events;
CREATE TABLE events
(
    id         INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    timestamp  DATETIME,
    event_type INTEGER,
    quantity   DECIMAL(7, 2) DEFAULT NULL,
    modified   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);