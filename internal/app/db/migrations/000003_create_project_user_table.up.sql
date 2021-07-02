CREATE TABLE `projects_user`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `project_id` int(11) UNSIGNED NOT NULL,
    `user_id` int(11) UNSIGNED NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_projects FOREIGN KEY(`project_id`) REFERENCES projects(`id`),
    CONSTRAINT FK_users FOREIGN KEY(`user_id`) REFERENCES users(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;