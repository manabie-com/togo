
package com.manabie.todotaskapplication.config;

import org.springframework.cache.annotation.CachingConfigurerSupport;
import org.springframework.cache.interceptor.KeyGenerator;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.util.StringUtils;

import java.lang.reflect.Method;

@Configuration
public class CacheConfig extends CachingConfigurerSupport {

    private static final String KEY_DELIMITER = "_";

    @Bean
    @Override
    public KeyGenerator keyGenerator() {
        return (Object target, Method method,
                Object... params) -> new StringBuilder().append(target.getClass().getSimpleName()).append(KEY_DELIMITER)
                .append(method.getName()).append(KEY_DELIMITER)
                .append(StringUtils.arrayToDelimitedString(params, KEY_DELIMITER));
    }

}