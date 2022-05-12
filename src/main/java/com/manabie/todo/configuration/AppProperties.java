package com.manabie.todo.configuration;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

@Data
@ConfigurationProperties(prefix = "app")
public class AppProperties {
  private Long defaultQuoteTask;
  private String jwtSecretKey;
  private String jwtIssuer;
  private Long jwtExpireInSeconds;
}
