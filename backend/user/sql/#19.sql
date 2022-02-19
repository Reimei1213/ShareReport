CREATE DATABASE IF NOT EXISTS share_report;

USE share_report;

CREATE TABLE IF NOT EXISTS user (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL,
  valid BOOLEAN NOT NULL DEFAULT 1,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `group` (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  valid BOOLEAN NOT NULL DEFAULT 1,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS group_user (
  id INT NOT NULL AUTO_INCREMENT,
  user_id VARCHAR(36) NOT NULL,
  group_id INT NOT NULL,
  valid BOOLEAN NOT NULL DEFAULT 1,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id),
  INDEX user_id_index(user_id),
  FOREIGN KEY (user_id)
    REFERENCES user(id),
  INDEX group_id_index(group_id),
  FOREIGN KEY (group_id)
    REFERENCES `group`(id)
);