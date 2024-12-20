CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  nick VARCHAR(50) NOT NULL unique,
  email VARCHAR(50) NOT NULL unique,
  password VARCHAR(255) NOT NULL,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP()  
);

CREATE TABLE followers(
  userId INTEGER not null,
  followerId INTEGER not null,

  PRIMARY KEY(userId, followerId),
  
  constraint fk_userId FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE,
  constraint fk_followerId FOREIGN KEY (followerId) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE posts(
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) not null,
  content varchar(300) not null,
  authorId int not null,
  likes int default 0,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (authorId) REFERENCES users(id) ON DELETE CASCADE
)