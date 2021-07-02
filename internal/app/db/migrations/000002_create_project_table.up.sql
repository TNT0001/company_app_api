CREATE TABLE `projects` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `category` ENUM("client","non-billable","system") NOT NULL DEFAULT "client",
    `projected_spend` int NOT NULL DEFAULT 0,
    `projected_variance` int NOT NULL DEFAULT 0,
    `revenue_recognised` int NOT NULL DEFAULT 0,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;