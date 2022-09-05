CREATE TABLE utilizatori (
  nume text PRIMARY KEY CHECK (nume NOT LIKE '% %'),
  email text UNIQUE,
  parola text NOT NULL,
  rol text NOT NULL DEFAULT 'utilizator',
  data_creatie timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE probleme (
  nume text PRIMARY KEY,
  titlu text NOT NULL,
  descriere text NOT NULL,
  dificultate integer NOT NULL,
  limita_timp integer NOT NULL,
  limita_memorie integer NOT NULL,
  evaluator text NOT NULL,
  autor text NOT NULL REFERENCES utilizatori(nume) ON DELETE CASCADE
);

CREATE TABLE solutii (
  id serial PRIMARY KEY,
  sursa text NOT NULL,
  problema text NOT NULL REFERENCES probleme(nume) ON DELETE CASCADE,
  utilizator text NOT NULL REFERENCES utilizatori(nume) ON DELETE CASCADE
);

CREATE TABLE teste (
  id serial PRIMARY KEY,
  intrare text NOT NULL,
  iesire text NOT NULL,
  problema text NOT NULL REFERENCES probleme(nume) ON DELETE CASCADE
);

CREATE TABLE rezultate (
  id serial PRIMARY KEY,
  timp integer,
  memorie integer,
  solutie integer NOT NULL REFERENCES solutii(id) ON DELETE CASCADE,
  test integer NOT NULL REFERENCES teste(id) ON DELETE CASCADE
);
