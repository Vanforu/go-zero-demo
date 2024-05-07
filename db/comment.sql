CREATE TABLE IF NOT EXISTS `comment` (
	`id` varchar(128) NOT NULL DEFAULT '' COMMENT '评论ID',
	`pid` varchar(128) NULL DEFAULT '' COMMENT '回复的评论ID',
	`aid` varchar(128) NULL DEFAULT '' COMMENT '评论所属文章',
	`content` varchar(1024) NULL DEFAULT '' COMMENT '评论内容',
	`creator` varchar(128) NULL COMMENT '评论人',
	`createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
