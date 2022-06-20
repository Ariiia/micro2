CREATE TABLE IF NOT EXISTS visits (counter INTEGER);
INSERT INTO visits VALUES (0);


CREATE TABLE IF NOT EXISTS notes (
id SERIAL,
body text
);

INSERT INTO notes VALUES (1, 'Note1 ... Once upon a time');
INSERT INTO notes VALUES (2, 'Note2 ... When I was older');
INSERT INTO notes VALUES (3, 'Note3 ... Execute my vision');
