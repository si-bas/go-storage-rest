-- +goose Up
-- +goose StatementBegin
CREATE TABLE `clients` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `code` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `key` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT current_timestamp(),
    `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`code`, `name`, `key`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE `clients`;

-- +goose StatementEnd