/*
 * Copyright (c) 2022, 2022 manabie.com and/or its affiliates. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.manabie.todo.configuration;

import com.google.common.collect.Lists;
import io.swagger.annotations.SwaggerDefinition;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurationSupport;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.OAuthBuilder;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.AuthorizationScope;
import springfox.documentation.service.Contact;
import springfox.documentation.service.SecurityScheme;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import java.io.IOException;
import java.io.InputStream;
import java.util.*;

import static springfox.documentation.builders.PathSelectors.regex;

/**
 *
 * @author longtran
 * @version 1.0
 */
@Configuration
@EnableSwagger2
public class SwaggerConfiguration extends WebMvcConfigurationSupport {

    private final Set<String> defaultProducesAndConsumes =
            new HashSet<>(Arrays.asList("application/json"));

    private List<AuthorizationScope> scopes() {
        return Lists.newArrayList(
                new AuthorizationScope("write", "write and read"),
                new AuthorizationScope("read", "read only"));
    }

    @Bean
    public SecurityScheme oauth() {
        return new OAuthBuilder()
                .name(SwaggerDefinition.Scheme.HTTPS.name())
                .scopes(scopes())
                .build();
    }

    @Bean
    public Docket productApi() {
        return new Docket(DocumentationType.SWAGGER_2)
                .select()
                .apis(RequestHandlerSelectors.basePackage("com.manabie.todo.endpoint"))
                .paths(regex("/api.*"))
                .build()
                .apiInfo(metaData())
                .consumes(defaultProducesAndConsumes)
                .produces(defaultProducesAndConsumes)
                .securitySchemes(Lists.newArrayList(oauth()))
                .enable(true);
    }

    /**
     * @return
     */
    private Properties getManifestProperties() {
        Properties result = new Properties();
        try (InputStream stream = getClass().getClassLoader().getResourceAsStream("META-INF/MANIFEST.MF")) {
            if (stream != null) {
                result.load(stream);
            } else {
                //LOG.error("File [META-INF/MANIFEST.MF] was not found");
            }
        } catch (IOException e) {
            //LOG.error("Cannot read manifest file properties", e);
        }

        return result;
    }

    /**
     * @return
     */
    public String getManifestBuildNumber() {
        String result = "Build #API-1.0.0-release-398a0145-cc27-44fa-81e7-094953066823";
        Properties allProps = getManifestProperties();
        //LOG.debug("Manifest file properties are [{}]", allProps);
        if (allProps != null && allProps.containsKey("Build-Revision") && allProps.containsKey("Build-Timestamp")) {
            result = String.format("Build #API-%s-release-%s, build on %s",
                    allProps.getProperty("Specification-Version"),
                    allProps.getProperty("Build-Revision"),
                    allProps.getProperty("Build-Timestamp"));
        }

        //LOG.debug("Implementation version is [{}]", result);
        return result;
    }

    private ApiInfo metaData() {
        return new ApiInfoBuilder()
                .title("Manabie Service API")
                .description("Manabie (@Manabie)")
                .version(getManifestBuildNumber())
                .license("Apache License Version 2.0")
                .licenseUrl("https://www.apache.org/licenses/LICENSE-2.0\"")
                .contact(new Contact("Contact the developer", "/about/", "Manabie"))
                .build();
    }

    @Override
    public void addResourceHandlers(final ResourceHandlerRegistry registry) {
        // Make Swagger meta-data available via <baseURL>/v2/api-docs/
        registry.addResourceHandler("/webjars/**").addResourceLocations("classpath:/META-INF/resources/webjars/");
        // Make Swagger UI available via <baseURL>/swagger-ui.html
        registry.addResourceHandler("/**").addResourceLocations("classpath:/META-INF/resources/");
    }
}
