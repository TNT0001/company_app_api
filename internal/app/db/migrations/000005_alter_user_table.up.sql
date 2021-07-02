ALTER TABLE `users`
ADD COLUMN `company_id` int NOT NULL AFTER `id`;

ALTER TABLE `users`
ADD CONSTRAINT FK_Companies
FOREIGN KEY (`company_id`) REFERENCES companies(`id`); 