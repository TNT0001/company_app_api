ALTER TABLE `users`
DROP FOREIGN KEY FK_Companies;

ALTER TABLE `users`
DROP COLUMN `company_id`;