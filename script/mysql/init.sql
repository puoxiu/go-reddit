CREATE DATABASE IF NOT EXISTS reddit;

USE reddit;

CREATE TABLE `users` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) Collate utf8mb4_general_ci NOT NULL,
    `password` varchar(64) Collate utf8mb4_general_ci NOT NULL,
    `email` varchar(64) Collate utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `community_id` int(10)  NOT NULL,
    `community_name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
    VALUES (1, 'Golang', 'A community about Golang.');
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
    VALUES (2, 'C++', 'A community about C++.');
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
    VALUES (3, '数据结构与算法', 'A community about alto.');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子ID',
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '帖子标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '帖子内容',
    `author_id` bigint(20) NOT NULL COMMENT '作者ID',
    `community_id` int(11) NOT NULL COMMENT '所属社区ID',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '帖子状态 0-正常 1-删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


