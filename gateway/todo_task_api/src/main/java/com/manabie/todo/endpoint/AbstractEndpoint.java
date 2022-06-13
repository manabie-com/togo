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

package com.manabie.todo.endpoint;

import com.manabie.todo.model.ErrorDetails;
import com.manabie.todo.model.APIResponseEntity;
import io.reactivex.Observable;
import io.reactivex.Observer;
import io.reactivex.disposables.Disposable;
import io.reactivex.schedulers.Schedulers;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.MessageSource;
import org.springframework.context.NoSuchMessageException;
import org.springframework.core.env.Environment;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.context.request.async.DeferredResult;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.Locale;
import java.util.stream.Collectors;

/**
 * @author longtran
 * @version 1.0
 */
@CrossOrigin(origins = "*", allowedHeaders = "*")
public abstract class AbstractEndpoint {

    protected static final String EMPTY = "";
    public final String TOKEN_PREFIX = "Bearer";

    @Autowired
    private MessageSource messageSource;

    @Autowired
    private Environment environment;

    @GetMapping("/ping")
    public @ResponseBody
    ResponseEntity<String> get() {
        return new ResponseEntity<String>("OK", HttpStatus.OK);
    }

    /**
     * Get massage localization.
     *
     * @param message A variable of type String.
     * @param locale  variable of type Locate.
     * @return
     */
    private String getMessageLocalization(String message, Locale locale) {
        try {
            return messageSource.getMessage(message.replaceAll("[\\{|}|']", ""), null, locale);
        } catch (NoSuchMessageException noSuchMessageException) {
            return message;
        }
    }

    /**
     * ResolveBindingResultErrors.
     *
     * @param bindingResult A variable of type BindingResult.
     * @return
     */
    private String resolveBindingResultErrors(BindingResult bindingResult, Locale locale) {
        return bindingResult.getFieldErrors().stream()
                .map(fr -> {
                    String validationMessage = StringUtils.isBlank(fr.getDefaultMessage()) ? "" : fr.getDefaultMessage();
                    return getMessageLocalization(validationMessage, locale);
                }).findFirst().orElse(bindingResult.getAllErrors().stream().map(x -> {
                    String validationMessage = StringUtils.isBlank(x.getDefaultMessage()) ? "" : x.getDefaultMessage();
                    return getMessageLocalization(validationMessage, locale);
                }).collect(Collectors.joining(", ")));
    }

    /**
     * convertInstanceOfObject.
     *
     * @return
     */
    private <T> T convertInstanceOfObject(Object o, Class<T> clazz) {
        try {
            return clazz.cast(o);
        } catch (ClassCastException e) {
            return null;
        }
    }

    /**
     * DeferredResult.
     *
     * @param details A variable of type Observable.
     * @return
     */
    @SuppressWarnings("unchecked")
    protected <P, T, B extends BindingResult, L extends Locale> DeferredResult<P> toDeferredResult(DeferredResult<P> deferredResult,
                                                                                                   Observable<T> details,
                                                                                                   B bindingResult,
                                                                                                   L locale,
                                                                                                   HttpServletRequest httpServletRequest,
                                                                                                   HttpServletResponse httpServletResponse) {
        details.subscribe(new Observer<T>() {
            @Override
            public void onSubscribe(Disposable d) {

            }

            @Override
            public void onNext(T t) {
                ResponseEntity<APIResponseEntity> response = new ResponseEntity<>(
                        new APIResponseEntity<>(t, HttpStatus.OK), HttpStatus.OK);
                deferredResult.setResult((P) response);
            }

            @Override
            public void onError(Throwable error) {
                error.printStackTrace();
                String validationMessage = StringUtils.isBlank(error.getMessage()) ? "" : error.getMessage();
                ErrorDetails errorDetails = new ErrorDetails(HttpStatus.BAD_REQUEST.value(),
                        bindingResult.hasErrors() ? resolveBindingResultErrors(bindingResult, locale) :
                                getMessageLocalization(validationMessage, locale));
                ResponseEntity<ErrorDetails> response = new ResponseEntity<>(errorDetails, HttpStatus.BAD_REQUEST);
                deferredResult.setErrorResult(response);
            }

            @Override
            public void onComplete() {

            }
        });
        return deferredResult;
    }

    /**
     * DeferredResult.
     *
     * @param deferredResult A variable of type DeferredResult.
     * @param details        Observable.
     * @return
     */
    @SuppressWarnings("unchecked")
    protected <P, T, L extends Locale> DeferredResult<P> toDeferredResult(DeferredResult<P> deferredResult,
                                                                          Observable<T> details,
                                                                          L locale,
                                                                          HttpServletRequest httpServletRequest,
                                                                          HttpServletResponse httpServletResponse) {
        details.subscribe(new Observer<T>() {
            @Override
            public void onSubscribe(Disposable d) {

            }

            @Override
            public void onNext(T t) {
                ResponseEntity<APIResponseEntity> response = new ResponseEntity<>(
                        new APIResponseEntity<>(t, HttpStatus.OK), HttpStatus.OK);
                deferredResult.setResult((P) response);
            }

            @Override
            public void onError(Throwable error) {
                error.printStackTrace();
                String validationMessage = StringUtils.isBlank(error.getMessage()) ? "" : error.getMessage();
                ErrorDetails errorDetails = new ErrorDetails(HttpStatus.BAD_REQUEST.value(),
                        getMessageLocalization(validationMessage, locale));
                ResponseEntity<ErrorDetails> response = new ResponseEntity<>(errorDetails, HttpStatus.BAD_REQUEST);
                deferredResult.setErrorResult(response);
            }

            @Override
            public void onComplete() {
                deferredResult.onCompletion(() -> {

                });
            }
        });
        return deferredResult;
    }

    /**
     * toObservable.
     *
     * @param bindingResult A variable of type BindingResult.
     * @return
     */
    protected <B extends BindingResult> Observable<B> toObservable(B bindingResult) {
        return Observable.fromCallable(() -> bindingResult)
                .flatMap(
                        v -> {
                            if (v.hasErrors()) {
                                throw new RuntimeException(resolveBindingResultErrors(v, Locale.getDefault()));
                            } else {
                                return Observable.just(v);
                            }
                        },
                        Observable::error,
                        Observable::empty
                ).subscribeOn(Schedulers.io());
    }

}
