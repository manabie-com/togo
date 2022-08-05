package com.example.server.config.database

import io.r2dbc.spi.ConnectionFactory
import org.springframework.boot.test.context.TestConfiguration
import org.springframework.context.annotation.Bean
import org.springframework.core.io.ClassPathResource
import org.springframework.r2dbc.connection.init.ConnectionFactoryInitializer
import org.springframework.r2dbc.connection.init.ResourceDatabasePopulator

@TestConfiguration
class DatabaseConfigTest {

    @Bean
    fun initializer(
        connectionFactory: ConnectionFactory,
    ): ConnectionFactoryInitializer = ConnectionFactoryInitializer().apply {
        setConnectionFactory(connectionFactory)
        setDatabasePopulator(
            ResourceDatabasePopulator(
                ClassPathResource("clean-database.sql"),
                ClassPathResource("schema.sql"),
                ClassPathResource("seed-database.sql"),
            )
        )
    }

}