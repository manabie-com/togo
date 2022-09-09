package com.manabie.todotaskapplication;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@SpringBootApplication
@EnableJpaRepositories
@EnableCaching
@ComponentScan(basePackages = {"com.manabie.todotaskapplication.*"})
public class TodoTaskApplication {

    public static void main(String[] args) {
        SpringApplication.run(TodoTaskApplication.class, args);
    }

}
