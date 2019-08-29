-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;

-- ---
-- Table 'user'
-- 
-- ---

DROP TABLE IF EXISTS `user`;
		
CREATE TABLE `user` (
  `id` INTEGER NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(20) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- ---
-- Table 'chat'
-- 
-- ---

DROP TABLE IF EXISTS `chat`;
		
CREATE TABLE `chat` (
  `id` INTEGER NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- ---
-- Table 'message'
-- 
-- ---

DROP TABLE IF EXISTS `message`;
		
CREATE TABLE `message` (
  `id` INTEGER NOT NULL AUTO_INCREMENT,
  `chat_id` INTEGER NOT NULL,
  `author_id` INTEGER NOT NULL,
  `text` MEDIUMTEXT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- ---
-- Table 'user_chat'
-- 
-- ---

DROP TABLE IF EXISTS `user_chat`;
		
CREATE TABLE `user_chat` (
  `user_id` INTEGER NOT NULL,
  `chat_id` INTEGER NOT NULL,
  PRIMARY KEY (`user_id`, `chat_id`)
);

-- ---
-- Foreign Keys 
-- ---

ALTER TABLE `message` ADD FOREIGN KEY (chat_id) REFERENCES `chat` (`id`);
ALTER TABLE `message` ADD FOREIGN KEY (author_id) REFERENCES `user` (`id`);
ALTER TABLE `user_chat` ADD FOREIGN KEY (user_id) REFERENCES `user` (`id`);
ALTER TABLE `user_chat` ADD FOREIGN KEY (chat_id) REFERENCES `chat` (`id`);

-- ---
-- Table Properties
-- ---

-- ALTER TABLE `user` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `chat` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `message` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `user_chat` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ---
-- Test Data
-- ---

-- INSERT INTO `user` (`id`,`username`,`created_at`) VALUES
-- ('','','');
-- INSERT INTO `chat` (`id`,`name`,`created_at`) VALUES
-- ('','','');
-- INSERT INTO `message` (`id`,`chat_id`,`author_id`,`text`,`created_at`) VALUES
-- ('','','','','');
-- INSERT INTO `user_chat` (`user_id`,`chat_id`) VALUES
-- ('','');