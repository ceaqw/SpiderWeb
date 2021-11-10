-- 项目初始化sql数据
-- 系统启动，项目会自动更新track的表,数据需要手动执行插入
-- member,role

-- member表,initUser密码nint123
insert into spider_web.member('name', 'password', 'role') values('nint', '$2a$10$JxXlceK7ZAYlGP0dfnEu3OL3mpnOCmRrMp1ksO9Xe6Hnbwpp8McyK', 1);

-- role表
insert into spider_web.role values (1, 'admin'), (2, 'customer');

-- crawler_information
CREATE TABLE `crawler_information` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project` varchar(200) CHARACTER SET latin1 NOT NULL COMMENT '项目',
  `platform` varchar(45) CHARACTER SET latin1 NOT NULL,
  `script_id` int(11) NOT NULL COMMENT '调用的lua的id',
  `ip` varchar(45) CHARACTER SET latin1 NOT NULL,
  `comment` text COMMENT '详细信息',
  `server` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `critical_kpi` varchar(200) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='爬虫脚本信息表'