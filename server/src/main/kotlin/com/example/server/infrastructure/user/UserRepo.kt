package com.example.server.infrastructure.user

import org.springframework.data.repository.kotlin.CoroutineCrudRepository
import org.springframework.stereotype.Repository
import java.util.*

@Repository
interface UserRepo: CoroutineCrudRepository<User, String> {

}