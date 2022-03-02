USE share_report;

ALTER TABLE `group` ADD user_id VARCHAR(36) NOT NULL AFTER id;
ALTER TABLE `group` ADD FOREIGN KEY (user_id) REFERENCES user(id);
ALTER TABLE `group` ADD INDEX user_id_index(user_id);