-- migrate:up
CREATE TABLE public."user" (
  id   INT PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);


-- migrate:down
DROP TABLE public."user"
