CREATE TABLE `articles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `body` varchar(255) NOT NULL,
  `position` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `threads` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `position` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `comments` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `thread_id` BIGINT NOT NULL,
  `body` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `thread_id_index` (`thread_id`)
);
