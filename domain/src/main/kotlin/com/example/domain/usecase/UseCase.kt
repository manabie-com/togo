package com.example.domain.usecase

import com.example.domain.util.suspendableResult
import com.example.domain.util.Result


abstract class UseCase<in Input : UseCase.Params, out Output : Any> {

    abstract suspend fun run(params: Input): Output

    suspend operator fun invoke(params: Input): Result<Output> = suspendableResult {
        params.selfValidate()?.let { error -> throw error }
        run(params)
    }

    abstract class Params {
        open val validators: List<Validator> = listOf()

        open fun selfValidate(): Error? {
            for (validator in validators) {
                when (val error = validator.validate()) {
                    null -> continue
                    else -> return error
                }
            }

            return null
        }
    }
}