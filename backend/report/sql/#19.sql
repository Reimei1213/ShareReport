CREATE DATABASE IF NOT EXISTS share_report;

USE share_report;

CREATE TABLE IF NOT EXISTS report (
  id INT NOT NULL AUTO_INCREMENT,
  book_id INT NOT NULL,
  group_user_id INT NOT NULL,
  content TEXT NOT NULL,
  star INT NOT NULL CHECK(1 <= star AND star <= 5),
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id),
  INDEX book_id_index(book_id),
  INDEX group_user_id_index(group_user_id)
);

CREATE TABLE IF NOT EXISTS tag (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS report_tag (
  id INT NOT NULL AUTO_INCREMENT,
  report_id INT NOT NULL,
  tag_id INT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id),
  INDEX report_id_index(report_id),
  FOREIGN KEY (report_id)
    REFERENCES report(id),
  INDEX tag_id_index(tag_id),
  FOREIGN KEY (tag_id)
    REFERENCES tag(id)
);