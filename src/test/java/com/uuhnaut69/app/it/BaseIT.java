package com.uuhnaut69.app.it;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.testcontainers.containers.PostgreSQLContainer;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;

/**
 * @author uuhnaut
 */
@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
@Testcontainers
@DirtiesContext
public class BaseIT {

  @Autowired
  protected TestRestTemplate testRestTemplate;

  @Container
  public static PostgreSQLContainer<?> postgresDB = new PostgreSQLContainer<>("postgres")
      .withDatabaseName("postgres")
      .withUsername("postgres")
      .withPassword("postgres")
      .withInitScript("init.sql");


  @DynamicPropertySource
  public static void properties(DynamicPropertyRegistry registry) {
    registry.add("spring.datasource.url", postgresDB::getJdbcUrl);
    registry.add("spring.datasource.username", postgresDB::getUsername);
    registry.add("spring.datasource.password", postgresDB::getPassword);
  }
}
