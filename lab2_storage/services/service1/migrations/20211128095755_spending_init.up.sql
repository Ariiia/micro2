CREATE TABLE IF NOT EXISTS visits (counter INTEGER);
INSERT INTO visits VALUES (0);


CREATE TABLE IF NOT EXISTS notes (
id SERIAL,
body text
);

INSERT INTO notes(body) VALUES ('Note1 ... Once upon a time');
INSERT INTO notes(body) VALUES ('Note2 ... When I was older');
INSERT INTO notes(body) VALUES ('Note3 ... Execute my vision');
