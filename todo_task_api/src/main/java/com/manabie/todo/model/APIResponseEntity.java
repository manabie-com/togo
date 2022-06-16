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
package com.manabie.todo.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.gson.annotations.SerializedName;
import io.swagger.annotations.ApiModelProperty;
import org.springframework.http.HttpStatus;
import org.springframework.lang.Nullable;
import org.springframework.util.Assert;

import java.io.Serializable;

/**
 *
 * @author longtran
 * @version 1.0
 */

public class APIResponseEntity<T> implements Serializable {

    @ApiModelProperty(example = "200")
    @JsonProperty("status")
    @SerializedName("status")
    public int status;

    @ApiModelProperty(example = "message")
    @JsonProperty("message")
    @SerializedName("message")
    public String message = "message";

    @ApiModelProperty(example = "entity", required = false)
    @JsonProperty("entity")
    @SerializedName("entity")
    @JsonInclude(JsonInclude.Include.NON_NULL)
    public T body;

    /**
     * Create a new {@code KleioAPIResponseEntity} with the given body and status code, and no headers.
     *
     * @param body   the entity body
     * @param status the status code
     */

    public APIResponseEntity(@Nullable T body, HttpStatus status) {
        Assert.notNull(status, "HttpStatus must not be null");
        Assert.notNull(body, "T body must not be null");
        this.status = status.value();
        this.body = body;
    }

    /**
     *
     * @param status the status code
     */
    public APIResponseEntity(String message, HttpStatus status) {
        Assert.notNull(message, "message must not be null");
        Assert.notNull(status, "HttpStatus must not be null");
        this.status = status.value();
        this.message = message;
    }

    public int getStatus() {
        return status;
    }

    public void setStatus(int status) {
        this.status = status;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public T getBody() {
        return body;
    }

    public void setBody(T body) {
        this.body = body;
    }
}
