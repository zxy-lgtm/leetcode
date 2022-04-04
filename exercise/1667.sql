-- Write your MySQL query statement below

-- mysql 字符串函数各种复习
-- CONCAT(str1, str2)：字符连接函数
-- UPPER(str)：将字符串改为大写字母
-- LOWER(str)：将字符串改为小写字母
-- LENGTH(str)：判定字符串长度
-- SUBSTRING(str, a, b):提取字段中的一段，从字符串str的第a位开始提取，提取b个字符
-- LEFT(str, n)：提取字符串最左边的n个字符
-- RIGHT(str, n)：提取字符串最右边的n个字符
SELECT a.user_id,
    CONCAT(UPPER(SUBSTRING(a.name, 1, 1)), LOWER(SUBSTRING(a.name, 2))) AS name
FROM Users a ORDER BY a.user_id;

/* 用left要快一点点*/