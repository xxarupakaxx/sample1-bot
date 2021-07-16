DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user(
  id   VARCHAR(255),
  displayName VARCHAR(255),
  id_type VARCHAR(255),
  timestamp TIMESTAMP NOT NULL ,
  reply_token VARCHAR(255)
);
