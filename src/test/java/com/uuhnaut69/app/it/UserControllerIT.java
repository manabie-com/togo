package com.uuhnaut69.app.it;

import com.uuhnaut69.app.user.model.User;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.springframework.http.HttpStatus;
import org.springframework.test.context.jdbc.Sql;

/**
 * @author uuhnaut
 */
public class UserControllerIT extends BaseIT {

  @Test
  @Order(1)
  @Sql({"/user.sql"})
  void findAllUser() {

    var response = testRestTemplate.getForEntity("/users", User[].class);
    var body = response.getBody();

    Assertions.assertNotNull(body);
    Assertions.assertEquals(1, body.length);
  }

  @Test
  @Order(2)
  @Sql({"/user.sql"})
  void findByIdSuccess() {

    var response = testRestTemplate.getForEntity("/users/{userId}", User.class, 1);
    var body = response.getBody();

    Assertions.assertNotNull(body);
    Assertions.assertEquals(1L, body.getId());
  }

  @Test
  @Order(3)
  @Sql({"/user.sql"})
  void findByIdNotFound() {

    var response = testRestTemplate.getForEntity("/users/{userId}", User.class, 2);

    Assertions.assertEquals(HttpStatus.NOT_FOUND, response.getStatusCode());
  }
}
