CREATE TABLE IF NOT EXISTS `article` (
	`id` varchar(128) NOT NULL DEFAULT '' COMMENT '文章ID',
	`title` varchar(128) NOT NULL DEFAULT '' COMMENT '文章标题',
	`desc` varchar(256) NOT NULL DEFAULT '' COMMENT '文章描述',
    `content` TEXT NOT NULL COMMENT '文章内容',
	`author` varchar(128)NOT NULL COMMENT '文章作者',
	`createTime` bigint NOT NULL DEFAULT 0 COMMENT '文章发布时间',
	`assets` varchar(1024) NOT NULL COMMENT '附件',
	`status` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '文章状态: 0-正常、1-已删除、99-草稿',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
