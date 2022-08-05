package com.example.domain.usecase.user

import com.example.domain.model.UserEntity

interface UserDateSource {
    suspend fun getUser(
        userId: String,
    ): UserEntity?
}