package com.manabie.todo;

import com.manabie.todo.entity.Todo;
import com.manabie.todo.entity.UserTodo;
import com.manabie.todo.repository.TodoRepository;
import com.manabie.todo.repository.UserTodoRepository;
import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.List;

@RunWith(SpringRunner.class)
@SpringBootTest
public class TodoServiceUnitTest {

    @Autowired
    private TodoRepository todoRepository;

    @Autowired
    private UserTodoRepository userTodoRepository;

    @Test
    public void whenInitializedDataMaximumTaskPerDay_thenFinds() {
        UserTodo userTodo = new UserTodo();
        userTodo.setUserId("f667cfa2-5b51-4266-afde-8a5963ec7d2a");
        userTodo.setMaximumTaskPerDay(1);
        userTodoRepository.create(userTodo);
        List<UserTodo> listUserTodos = userTodoRepository.getAll();
        Assert.assertNotEquals(0, listUserTodos.size());

        Todo todo = new Todo();
        todo.setDateCreated(System.currentTimeMillis());
        todo.setCreateByUserId(userTodo.getUserId());
        todo.setTitle("Unit test maximum task per day");
        todo.setName("Task 1");
        todo.setNote("Unit test maximum task per day");
        todoRepository.create(todo);

        Assert.assertEquals(1, todoRepository.getTaskBy(userTodo.getUserId()).size());
    }

}

