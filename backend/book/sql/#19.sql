CREATE DATABASE IF NOT EXISTS share_report;

USE share_report;

CREATE TABLE IF NOT EXISTS book (
  id INT NOT NULL AUTO_INCREMENT,
  isbn INT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id)
  INDEX isbn_index(isbn)
);