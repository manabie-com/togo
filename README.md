### Overview:
- User có thể tạo mới task. Mỗi ngày user được phép tạo tối đa N task
- Nếu user đã tạo max số task được tạo trong ngày. trả về lỗi 4xx cho client và từ chối request tạo mới task đó.

### Create table
```sql
-- users definition

CREATE TABLE `test`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  `max_todo` INT NULL DEFAULT 5,
  PRIMARY KEY (`id`));


-- tasks definition

CREATE TABLE `test`.`task` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `user_id` INT NOT NULL,
  `content` VARCHAR(45) NOT NULL,
  `created_date` INT NOT NULL,
  PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`)
    REFERENCES `test`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);
```

### Run
```cosole
cd togo
python3 app.py
```
### Call API
#### User register
```cosole
curl -X POST http://127.0.0.1:8000/signup -H 'content-type: application/json' -d '{ "user_name": "quan", "password": "1234"}'
```
#### Login
```cosole
curl -X POST http://127.0.0.1:8000/login -H 'cache-control: no-cache' -H 'content-type: application/json'-d '{"username": "quan", "password": "1234"}'
```
#### Create task
```cosole
curl -X POST http://127.0.0.1:8000/task -H 'authorization: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Im5hbSIsInBhc3N3b3JkIjoiYWFhYWEifQ.f12_9TkDfd0Pxv1lH6MngaCxkZgsmQBS_oIVDE9m_dw' -H 'content-type: application/json' -d '{"name": "tast1", "content": "content1"}'
```
#### Get list task
```cosole
curl -X GET http://127.0.0.1:8000/task -H 'authorization: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Im5hbSIsInBhc3N3b3JkIjoiYWFhYWEifQ.f12_9TkDfd0Pxv1lH6MngaCxkZgsmQBS_oIVDE9m_dw' -H 'content-type: application/json'
```
#### Check healthy
```cosole
curl -X GET http://127.0.0.1:8000/ping
```
### Run test
#### Unit test
```cosole
python3 -m unittest test/unit_test/user_test.py
python3 -m unittest test/unit_test/task_test.py
```
#### Integration test
```cosole
python3 -m unittest test/integration_test/user_test.py
python3 -m unittest test/integration_test/user_test.py
```
### About my solution
- Code theo mô hình MVC giúp dễ quản lí source và mở rộng các tính năng mới.
- Database mình chọn MySQL vì 2 lí do:
  - Dữ liệu đầu vào(user, task) có cấu trúc rõ ràng. Nếu không kiếm soát được dữ liệu đầu vào thì sẽ chọn NoSQL.
  - SQL giúp thể hiện được mối quan hệ giữa 2 bảng dữ liệu.
- Ngôn ngữ Python giúp việc dựng API rất nhanh và đơn giản.
### Extension
- Phân trang kết quả khi lấy danh sách task.
- Thêm các đầu api như sửa task, sửa thông tin user...
- Sử dụng redis cache để làm bộ đệm lưu thông tin user để giảm sô lượng truy vấn database và tăng tốc độ API.
- Thêm 1 số luật cho mật khẩu người dùng nếu cần độ bảo mật cao hơn.