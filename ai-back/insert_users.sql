-- 假设您已经创建了一个名为 users 的表，并且已经配置好了数据库连接

-- 插入第一条用户数据


INSERT INTO users (created_at,updated_at, name, password, email, last_login_time, chat_num, role, account_status)
VALUES ('2024-04-06 09:00:00','2024-04-06 09:00:00','王五', 'password3', '123@qq.com', '2024-04-06 09:00:00', 0, 'User', 'normal');

-- 第一条用户数据
INSERT INTO users (created_at, updated_at, name, password, email, last_login_time, chat_num, role, account_status)
VALUES ('2024-04-06 09:00:00', '2024-04-06 09:00:00', '王五', 'password3', 'wangwu@moonshotai.com', '2024-04-06 09:00:00', 0, 'User', 'normal');

-- 第二条用户数据
INSERT INTO users (created_at, updated_at, name, password, email, last_login_time, chat_num, role, account_status)
VALUES ('2024-04-06 10:00:00', '2024-04-06 10:00:00', '李四', 'password4', 'lisi@moonshotai.com', '2024-04-06 10:00:00', 0, 'User', 'normal');

-- 第三条用户数据
INSERT INTO users (created_at, updated_at, name, password, email, last_login_time, chat_num, role, account_status)
VALUES ('2024-04-06 11:00:00', '2024-04-06 11:00:00', '张三', 'password5', 'zhangsan@moonshotai.com', '2024-04-06 11:00:00', 0, 'User', 'normal');

-- 可以根据需要继续添加更多的INSERT语句

-- 例如：
-- INSERT INTO users ... VALUES ...;
-- INSERT INTO users ... VALUES ...;
-- ...