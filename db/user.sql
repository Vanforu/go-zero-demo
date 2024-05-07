CREATE TABLE IF NOT EXISTS `user` (
	`id` varchar(128) NOT NULL DEFAULT '' COMMENT '用户ID',
	`name` varchar(32) NOT NULL COMMENT '用户姓名',
	`phone` varchar(128) NOT NULL DEFAULT '' COMMENT '用户手机号/登录用户名',
	`password` varchar(128) NOT NULL DEFAULT '' COMMENT '用户密码',
	`type` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '用户类型: 0-普通用户、1-后台管理员',
	`createTime` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户注册时间',
	`status` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '用户状态: 0-正常、1-冻结、2-注销、3-拉黑',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
