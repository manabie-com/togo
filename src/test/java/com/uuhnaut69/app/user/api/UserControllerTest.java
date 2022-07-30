package com.uuhnaut69.app.user.api;

import static org.hamcrest.Matchers.hasSize;
import static org.hamcrest.Matchers.is;

import com.uuhnaut69.app.common.exception.NotFoundException;
import com.uuhnaut69.app.user.model.User;
import com.uuhnaut69.app.user.service.UserService;
import java.time.Instant;
import java.util.List;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

/**
 * @author uuhnaut
 */
@RunWith(SpringRunner.class)
@WebMvcTest(UserController.class)
class UserControllerTest {

  @Autowired
  private MockMvc mvc;

  @MockBean
  private UserService userService;

  @Test
  void findAllUser() throws Exception {
    var users = List.of(new User(1L, "uuhnaut69", 10, Instant.now()),
        new User(2L, "uuhnaut96", 69, Instant.now()));

    Mockito.when(userService.findAllUser()).thenReturn(users);

    mvc.perform(MockMvcRequestBuilders.get("/users"))
        .andExpect(MockMvcResultMatchers.status().isOk())
        .andExpect(MockMvcResultMatchers.jsonPath("$", hasSize(2)));
  }

  @Test
  void findUserDetailShouldSuccess() throws Exception {
    var user = new User(1L, "uuhnaut69", 10, Instant.now());

    Mockito.when(userService.findUserById(1L)).thenReturn(user);

    mvc.perform(MockMvcRequestBuilders.get("/users/{userId}", 1L))
        .andExpect(MockMvcResultMatchers.status().isOk())
        .andExpect(MockMvcResultMatchers.jsonPath("$.id", is(1)))
        .andExpect(MockMvcResultMatchers.jsonPath("$.limitConfig", is(10)));
  }

  @Test
  void findUserDetailShouldFailed() throws Exception {
    Mockito.when(userService.findUserById(1L)).thenThrow(new NotFoundException());

    mvc.perform(MockMvcRequestBuilders.get("/users/{userId}", 1L))
        .andExpect(MockMvcResultMatchers.status().isNotFound());
  }
}