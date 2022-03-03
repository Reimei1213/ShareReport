USE share_report;

RENAME TABLE `group` TO organization;

RENAME TABLE `group_user` TO organization_user;
ALTER TABLE organization_user CHANGE COLUMN group_id organization_id INT NOT NULL;
ALTER TABLE organization_user DROP INDEX group_id_index;
CREATE INDEX organization_id_index on organization_user(organization_id);