-- Clean up
DELETE
FROM todo;

INSERT INTO todo(task, user_id)
VALUES ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1),
       ('Test', 1);