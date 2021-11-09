-- 项目初始化sql数据
-- 系统启动，项目会自动更新track的表,数据需要手动执行插入
-- member,role

-- member表,initUser密码nint123
insert into spider_web.member('name', 'password', 'role') values('nint', '$2a$10$JxXlceK7ZAYlGP0dfnEu3OL3mpnOCmRrMp1ksO9Xe6Hnbwpp8McyK', 1);

-- role表
insert into spider_web.role values (1, 'admin'), (2, 'customer');