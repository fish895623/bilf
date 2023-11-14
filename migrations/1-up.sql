create table daily (
  id serial,
  usage integer,
  tagList integer[]
  PRIMARY KEY(id)
);

create table taglist (
  id serial,
  name text,
  PRIMARY KEY(id)
);
