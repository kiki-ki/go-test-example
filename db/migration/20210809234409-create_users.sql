
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `age` int NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;
