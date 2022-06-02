-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(100) NULL,
  `username` varchar(100) NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`),
  UNIQUE KEY `users_username_unique` (`username`),
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

DROP TABLE `users`;

-- +goose StatementEnd
