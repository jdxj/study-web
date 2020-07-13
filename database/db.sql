-- user_agent 用于记录 http request header 中的 User-Agent 字段
CREATE TABLE `user_agent` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `value` varchar(1024) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;