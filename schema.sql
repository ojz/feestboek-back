CREATE TABLE users (
    id          INTEGER PRIMARY KEY,
    username    TINYTEXT,
    hash        TINYTEXT,

    invitation  INTEGER REFERENCES invitations (id) ON DELETE CASCADE,
    bio         TINYTEXT
);

CREATE TABLE invitations (
    id          INTEGER PRIMARY KEY,
    userid      INTEGER REFERENCES users (id) ON DELETE CASCADE,
    code        TINYTEXT,
    used        BOOLEAN DEFAULT 0
);

INSERT INTO users (username, hash, bio) VALUES ("root", "x", "the start of it all");
INSERT INTO invitations (userid, code) VALUES (1, "yes");
INSERT INTO invitations (userid, code) VALUES (1, "welcome");
INSERT INTO invitations (userid, code) VALUES (1, "please enter");
