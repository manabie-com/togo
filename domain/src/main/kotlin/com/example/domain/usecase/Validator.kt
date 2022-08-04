package com.example.domain.usecase

abstract class Validator {
    companion object {
        const val CODE_MAX_LENGTH = 6
        const val CODE_MIN_LENGTH = 2
    }

    open fun validate(): Error? {
        return null
    }
}