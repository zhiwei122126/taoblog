-- 创建表 options
CREATE TABLE IF NOT EXISTS `options` (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `value` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    PRIMARY KEY(`id`)
);

-- 创建表 posts
CREATE TABLE IF NOT EXISTS `posts` (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `date` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    `modified` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    `title` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `slug` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `type` VARCHAR(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `taxonomy` INT(20) UNSIGNED NOT NULL DEFAULT 1,
    `status` ENUM('public', 'draft'),
    `page_view` INT(20) UNSIGNED NOT NULL DEFAULT 0,
    `comment_status` INT(1) UNSIGNED DEFAULT 1,
    `comments` INT(20) UNSIGNED NOT NULL DEFAULT 0,
    `metas` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `source` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `source_type` ENUM('html', 'markdown'),
    PRIMARY KEY(`id`)
);

-- 创建表 comments
CREATE TABLE IF NOT EXISTS `comments` (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT(20) UNSIGNED NOT NULL,
    `author` TINYTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `email` VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `url` VARCHAR(200) CHARACTER SET utf8 COLLATE utf8_general_ci,
    `ip` VARCHAR(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `date` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    `content` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `parent` INT(20) UNSIGNED NOT NULL,
    `ancestor` INT(20) UNSIGNED NOT NULL,
    PRIMARY KEY(`id`)
);

-- 创建表 taxonomies
CREATE TABLE IF NOT EXISTS taxonomies (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `slug` VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `parent` INT(20) UNSIGNED NOT NULL,
    `ancestor` INT(20) UNSIGNED NOT NULL,
    PRIMARY KEY(`id`)
);

-- 创建表 /分类/文章样式和JavaScript/关键字表 post_metas
CREATE TABLE IF NOT EXISTS post_metas (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `type` ENUM('post','page','tax'),
    `tid` INT(20) UNSIGNED NOT NULL,
    `header` TEXT CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `footer` TEXT CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `keywords` TEXT CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    PRIMARY KEY(`id`)
);

-- 创建表 tag标签/post_tags
CREATE TABLE IF NOT EXISTS tags (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `alias` INT(20) UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY(`id`),
    UNIQUE KEY(`name`)
);

-- 创建表 文章标签表
CREATE TABLE IF NOT EXISTS post_tags (
    `id` INT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT(20) UNSIGNED NOT NULL,
    `tag_id` INT(20) UNSIGNED NOT NULL,
    PRIMARY KEY(`id`)
);
