# database
名称 描述 价格
添加商品 查询商品列表（可控地根据价格排序，小数） 删除商品 更新价格

数据库
echo
gorm
go mod


任务：
建一张表，表里包含三个变量（商品名称 价格和描述）
在此基础上用户可以通过链接添加商品+查询商品（小数and排序）+删除商品+更新价格；这些用echo来实现

# mvc architecture
## Controller:
Handle all the request from the customer
Ask the model for data based on the request

## Model:
Handle data logic
Interacts with the database (adding, deleting, etc)
Only use the model to perform these interactions

## View:
How to present the information the controller sends it 
There will be a template file that dynamically rendered html

Model and view never interacts with each other

# SQL syntax
## Data definition language: 定义、操作数据结构（eg tables）
1. CREATE TABLE obj_name (column_name data_type)
Eg. CREATE sales (purchase_number INT);
2. ALTER (altering existing obj)
ALTER Table sale
ADD COLUMN date_purchase DATE;
3. DROP 
DROP TABLE customers
4. RENAME
RENAME TABLE customers TO users
5. Truncate (留有框架，but empty data）
Truncate Table customers

## Data manipulation language: 操作数据
1. SELECT * FROM (table里的所有东西）
2. SELECT … FROM 
    1. 还可以选择没有——的：SELECT * FROM sales WHERE not date_purchase = ‘2018’;
    2. SELECT COUNT(value) FROM sales WHERE condition; 出value的个数
    3. SELECT AVG(value) FROM sales WHERE condition; 出value的平均值
    4. SELECT SUM(value) FROM sales WHERE condition; 出value的总和
    5. SELECT MAX / MIN(value) FROM sales; 出value的最值
    6. SELECT COUNT(id) FROM sales GROUP BY value; 将id按value分组，同时得出每组有几个
        1. SELECT id SUM(value) FROM sales GROUP BY id having sum(value) > 400;
        2. SELECT count(id), value FROM sales GROUP BY id ORDER BY value (desc); 根据value从小到大排序，加desc从大到小
    7. SELECT id FROM sale where value IS (not) NULL; 
    8. SELECT id, value, date_purchase FROM sales WHERE date_purchase IN (‘2017’, ‘2018’, ‘2020’)；多选
    9. SELECT id, value, date_purchase FROM sales WHERE value BETWEEN 200 AND 500；范围选择
    10. SELECT id AS sales_id FROM sale; 把id出来变为sales_id
    11. SELECT value, date_purchase FROM sales AS sales; 把value和date变成表格展示出
3. INSERT INTO sales (purchase_num, data_purchase) Values (?,?)
4. UPDATE (修改）
UPDATE sales SET purchase_date = ‘2018’ WHERE purchase_num = 1;
5. DELETE FROM …（相比于truncate精准删除）
DELETE FROM sales WHERE purchase_num = 1;

## Data Control language: 给特定user准许
These rights will be assigned to a person who has a username registered at the Local Server (‘Localhost’: IP 127.0.0.1)
1. CREATE USER ‘delicatesse@localhost’ IDENTIFIED BY ‘password’
2. Grant
Eg. 只grant select
GRANT SELECT ON sale.customers to ‘delicatesse@localhost’ 
Grant all
GRANT ALL ON sale.* to ‘delicatesse@localhost’ 
3. REVOKE
REVOKE SELECT ON sale.customers FROM ‘delicatesse@localhost’ 

## Transaction control language: 帮助数据库保存
1. COMMIT
Save the changes you make 
Will let others have access to the database
Eg UPDATE customers SET last_name = ‘Quan’ WHERE customer_id = 7 COMMIT;
COMMIT了之后users才能看见改变

2. ROLLBACK
Allows you to undo any changes you made but don’t want to be saved permanently
Reverts to the last non-committed state
