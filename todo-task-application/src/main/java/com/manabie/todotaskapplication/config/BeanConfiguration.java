package com.manabie.todotaskapplication.config;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.module.SimpleModule;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder;
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter;
import org.springframework.web.client.RestTemplate;

import java.util.Arrays;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Configuration
public class BeanConfiguration {
    @Bean(name = "objectMapper")
    public ObjectMapper createObjectMapper() {
        ObjectMapper objectMapper = new ObjectMapper();
        objectMapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        return objectMapper;
    }

    @Bean
    public RestTemplate restTemplate() {
        ObjectMapper objectMapper = Jackson2ObjectMapperBuilder.json().build();
        SimpleModule module = new SimpleModule();
        objectMapper.registerModule(module);
        return new RestTemplate(Arrays.asList(new MappingJackson2HttpMessageConverter(objectMapper)));
    }


    @Bean
    public RestTemplate restTemplateRaw() {
        return new RestTemplate();
    }
}
