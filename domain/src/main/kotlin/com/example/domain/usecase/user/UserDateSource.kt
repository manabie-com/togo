package com.example.domain.usecase.user

import com.example.domain.model.UserEntity
import java.util.UUID

interface UserDateSource {
    suspend fun getUser(
        userId: String,
    ): UserEntity?
}