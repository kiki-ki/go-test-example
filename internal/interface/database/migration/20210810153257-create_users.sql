
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `age` int NOT NULL,
  UNIQUE KEY `uq_users_email` (`email`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;