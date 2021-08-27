package com.manabie.togo.test;
import com.manabie.togo.model.User;
import com.manabie.togo.services.UserService;
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
public class UserTest {

    @Autowired
    private UserService userService;
    
    @Test
    public void testCreateUser() throws Exception{
        String username = "thientp";
        String password = "deptraivl";
        long maxtodo = 5l;
        
        userService.addUser(username, password, maxtodo);
        User user = userService.loadUser(username);
        Assert.assertNotNull(user);        
    }
    
    @Test
    public void testLoadUser() throws Exception{
        String username = "thientp";
        String password = "deptraivl";
        long maxtodo = 5l;
        
        userService.addUser(username, password, maxtodo);
        User user = userService.loadUser(username);
        Assert.assertNotNull(user);            
    }
}
