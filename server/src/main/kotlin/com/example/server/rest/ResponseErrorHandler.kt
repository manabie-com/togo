package com.example.server.rest

import org.springframework.boot.autoconfigure.web.WebProperties.Resources
import org.springframework.boot.autoconfigure.web.reactive.error.AbstractErrorWebExceptionHandler
import org.springframework.boot.web.reactive.error.ErrorAttributes
import org.springframework.context.ApplicationContext
import org.springframework.core.annotation.Order
import org.springframework.http.HttpStatus
import org.springframework.http.codec.ServerCodecConfigurer
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.*
import com.example.server.rest.response.ErrorResponse


@Component
@Order(-2)
class ResponseErrorHandler(
    attributes: ErrorAttributes,
    resources: Resources,
    context: ApplicationContext,
    configurer: ServerCodecConfigurer,
) : AbstractErrorWebExceptionHandler(attributes, resources, context) {


    init {
        super.setMessageReaders(configurer.readers)
        super.setMessageWriters(configurer.writers)
    }

    override fun getRoutingFunction(attributes: ErrorAttributes) = coRouter {

        RequestPredicates.path("/api/**")
            .invoke { request ->
                val error = attributes.mapError(request)
                val response = when (error) {
                    else -> ErrorResponse(
                        statusCode = error.code,
                        error = error.message,
                    )
                }
                ServerResponse.status(HttpStatus.valueOf(error.code / 1000))
                    .json()
                    .bodyValueAndAwait(response)
            }

    }
}