-- +goose Up
-- +goose StatementBegin
CREATE TABLE `files` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `client_id` int(11) unsigned NOT NULL,
    `code` varchar(255) NOT NULL,
    `original_name` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `extension` varchar(255) NOT NULL,
    `size` bigint unsigned NOT NULL,
    `path` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT current_timestamp(),
    `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`code`),
    KEY `files_FK` (`client_id`),
    CONSTRAINT `files_FK` FOREIGN KEY (`client_id`) REFERENCES `clients` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE `files`;

-- +goose StatementEnd