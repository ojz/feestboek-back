CREATE TABLE users (
    id          INTEGER PRIMARY KEY,
    name        TINYTEXT,
    password    TINYTEXT,
    parent      INTEGER REFERENCES user (id) ON DELETE CASCADE
);

CREATE TABLE invitations (
    userid      INTEGER REFERENCES user (id) ON DELETE CASCADE,
    code        TINYTEXT
);
