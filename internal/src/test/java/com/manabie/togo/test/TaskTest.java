/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.manabie.togo.test;

import com.manabie.togo.model.Task;
import com.manabie.togo.model.User;
import com.manabie.togo.services.TaskService;
import com.manabie.togo.services.UserService;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.context.annotation.Bean;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootTest
public class TaskTest {
    @Autowired
    private TaskService taskService;
    
    private static final SimpleDateFormat formatter = new SimpleDateFormat("yyyy-MM-dd");
    
    @Test
    public void testListTasks() throws Exception{
        List<Task> list = taskService.listTasks();
        Assert.assertNotNull(list);
    }

    @Test
    public void testListTasksPerDay() throws Exception{
        Date today = new Date();
        String strToday = formatter.format(today);
        List<Task> list = taskService.listTaskInDay(strToday);
        Assert.assertNotNull(list);
    }

    @Test
    public void testListTasksUserPerDay() throws Exception{
        String username = "loda";
        Date today = new Date();
        String strToday = formatter.format(today);
        List<Task> list = taskService.listTaskPerDayUser(username, strToday);
        Assert.assertNotNull(list);
    }
}
