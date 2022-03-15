## Phân tích yêu cầu

### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.

**=> Keyword:**
- **One single API**
- **Todo task**: 
  - Có những loại todo task nào?
    - **Any.do** (free/premium options)
    - Google Tasks (miễn phí)
    - Todoist (free/premium options)
    - Evernote (free/premium options)
    - Wunderlist (miễn phí)
    - Microsoft To-Do (miễn phí)
    - Bear (free/premium options)
    - Ike (miễn phí)
    - ...
  - Những điểm chung của những todo task đó?
    - Checkbox
    - Content
    - Time
  - Những điểm riêng nào đáng học tập?
    - Treeview: nhóm các todo task theo category
    - Độ ưu tiên: sắp xếp hoặc nhóm các task theo độ ưu tiên
    - Todo task cho nhóm: cho các nhóm, tổ chức ...
    - Tích hợp comment để trao đổi trực tiếp về todo task
    - gắn cờ, gắn sao, gắn tag,...
    - Remind, duedate,...
    - cần đăng nhập, không cần đăng nhập
  - Bản thân có ý tưởng mới nào không?
    - Sau khi tham khảo những todo task khác thì thấy các chức năng đã rất đủ, mỗi todo app hướng tới 1 những đối tượng khác nhau. Do quá nhiều chức năng nên cảm thấy có những todo task đang đánh mất tính đơn giản của nó, bù lại là nhiều tính năng để lựa chọn 
    - Thêm bot cho chat để Remind user: Bình thường không phải ứng dụng nào cũng có mặt trên đa nền tảng, vì vậy nên add bot cho những phần mềm chat của user để dễ dàng remind cho user về todo task.

**=>** trong giới hạn về thời gian => sẽ tạo ra todo task với các điểm chung cơ bản

- **Max task/user/day**:
  - Phải có API tạo user, login
  - Tạo giới hạn task trong ngày
  - task trong ngày bằng sql query

- **Different Users-Different Limit**:
  -  Config max Task của user trong liên quan tới user

- **Records**:
  - Dùng PostgreSQL
