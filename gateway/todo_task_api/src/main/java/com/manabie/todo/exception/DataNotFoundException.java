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
package com.manabie.todo.exception;

import java.io.Serializable;

/**
 *
 * @author longtran
 * @version 1.0
 */
public class DataNotFoundException extends DataException {

    private static final long serialVersionUID = 2524686849100170713L;

    /**
     * Create a new <code>DataNotFoundException</code> with a message text and the root exception.
     *
     * @param message
     *            Message text as String
     * @param cause
     *            The root exception
     */
    public  DataNotFoundException(String message, Throwable cause) {
        super(message, cause);
    }

    /**
     * Create a new <code>DataNotFoundException</code> with a message text.
     *
     * @param message
     *            Message text as String
     */
    public  DataNotFoundException(String message) {
        super(message);
    }

    /**
     * Create a new <code>DataNotFoundException</code> with the id of the expected entity.
     *
     * @param id
     *            Id of the expected entity
     */
    public  DataNotFoundException(Serializable id) {
        super("Entity class not found in persistence layer, id: " + id);
    }
}
