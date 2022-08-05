package com.example.server.rest

import org.springframework.boot.web.reactive.error.ErrorAttributes
import org.springframework.http.HttpStatus
import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.server.ResponseStatusException
import com.example.domain.exception.Error


internal fun ErrorAttributes.mapError(request: ServerRequest) = when (val requestError = getError(request)) {
    is ResponseStatusException -> when (requestError.status) {
        HttpStatus.NOT_FOUND -> Error.NotFoundError(requestError)
        else -> Error.UnexpectedError(requestError)
    }
    is Error -> requestError

    else -> Error.UnexpectedError(requestError)

}