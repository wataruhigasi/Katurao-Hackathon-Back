CREATE TABLE `articles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `body` LONGTEXT NOT NULL,
  `position` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `threads` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `title` LONGTEXT NOT NULL,
  `position` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `keijiban_rakugaki` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `body` LONGTEXT NOT NULL,
  `position` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `keijiban_rakugaki_id_index` (`id`)
);

CREATE TABLE `thread_rakugaki` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `thread_id` BIGINT NOT NULL,
  `body` LONGTEXT NOT NULL,
  `position` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `thread_rakugaki_id_index` (`id`)
);
