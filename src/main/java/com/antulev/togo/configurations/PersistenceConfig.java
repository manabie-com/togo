package com.antulev.togo.configurations;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.domain.AuditorAware;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

@Configuration
@EnableJpaAuditing(auditorAwareRef = "auditorProvider")
public class PersistenceConfig {

	@Bean("auditorProvider")
    public AuditorAware<String> auditorProvider() {
        return new AuditorAwareImpl();
    }
}