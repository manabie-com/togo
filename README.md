### How to run your code locally?
uvicorn main:app
### A sample “curl” command to call your API
curl -X 'POST' \
  'http://127.0.0.1:8000/tasks/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "todo",
  "user_id": 1
}'
### How to run your unit tests locally?
pytest
### What do you love about your solution?
- Sử dụng python với framework FastAPI nên code khá là ít, đơn giản và tích hợp sẵn Swagger UI cho API documentation.  
- Database sử dụng sqlite có kích thước bé và portable.
- Sử dụng ORM SQLAlchemy nên có thể thay đổi database(Relational Database) dễ dàng mà không cần thay đổi code.