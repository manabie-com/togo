package com.example.domain.usecase

abstract class Validator {
    open fun validate(): Error? {
        return null
    }
}