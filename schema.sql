CREATE TABLE users (
    id          INTEGER PRIMARY KEY,
    username    TINYTEXT,
    password    TINYTEXT,
    parent      INTEGER REFERENCES users (id) ON DELETE CASCADE,
    bio         TINYTEXT
);

CREATE TABLE invitations (
    userid      INTEGER REFERENCES users (id) ON DELETE CASCADE,
    code        TINYTEXT
);
