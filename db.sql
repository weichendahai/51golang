CREATE TABLE articles (
  id Integer NOT NULL,
  title varchar,
  content text,
  created_at timestamp,
  CONSTRAINT [PK_articles] PRIMARY KEY ([id])
);
