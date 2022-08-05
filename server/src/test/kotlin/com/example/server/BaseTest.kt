package com.example.server

import com.example.domain.exception.Error
import com.example.server.infrastructure.task.TaskDS
import com.example.server.infrastructure.user.UserDS
import com.ninjasquad.springmockk.MockkBean
import io.mockk.coEvery
import kotlinx.coroutines.runBlocking
import okhttp3.mockwebserver.Dispatcher
import okhttp3.mockwebserver.MockResponse
import okhttp3.mockwebserver.MockWebServer
import okhttp3.mockwebserver.RecordedRequest
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.extension.ExtendWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.context.properties.EnableConfigurationProperties
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.ApplicationContext
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.test.context.ActiveProfiles
import org.springframework.test.context.TestPropertySource
import org.springframework.test.web.reactive.server.WebTestClient
import org.springframework.web.reactive.function.BodyInserters
import java.time.Duration
import java.util.*

@SpringBootTest
@ActiveProfiles("Test")
class BaseTest {
    companion object {
    }

    @Autowired
    lateinit var context: ApplicationContext

    lateinit var client: WebTestClient


    lateinit var clientUseKotlinSerialization: WebTestClient
    @Autowired
    lateinit var taskDS: TaskDS

    @Autowired
    lateinit var userDS: UserDS

    var server = MockWebServer()

    var ocServer = MockWebServer()


    @BeforeEach
    fun setUp() {
        client =
            WebTestClient.bindToApplicationContext(context).configureClient().responseTimeout(Duration.ofSeconds(1000))
                .build()
        clientUseKotlinSerialization =
            WebTestClient.bindToApplicationContext(context).configureClient().responseTimeout(Duration.ofSeconds(1000))
                .build()
        server.dispatcher = object : Dispatcher() {
            private val fallBackResponse = MockResponse()
                .setResponseCode(500)

            override fun dispatch(request: RecordedRequest): MockResponse = request.path?.let {
                return fallBackResponse
            } ?: fallBackResponse
        }
    }


    /**
     * Helper functions for asserting test result
     */
    protected fun WebTestClient.ResponseSpec.isJsonResponse(): WebTestClient.ResponseSpec = expectHeader()
        .contentType(MediaType.APPLICATION_JSON)

    protected fun WebTestClient.ResponseSpec.baseAssert() = isJsonResponse()
    protected fun WebTestClient.ResponseSpec.isOk(): WebTestClient.ResponseSpec = expectStatus().isOk
    protected fun WebTestClient.ResponseSpec.isNotFound(): WebTestClient.ResponseSpec = expectStatus().isNotFound
    protected fun WebTestClient.ResponseSpec.isUnauthorized(): WebTestClient.ResponseSpec =
        expectStatus().isUnauthorized

    protected fun WebTestClient.ResponseSpec.isBadRequest(): WebTestClient.ResponseSpec = expectStatus().isBadRequest
    protected fun WebTestClient.ResponseSpec.is5xxServerError(): WebTestClient.ResponseSpec =
        expectStatus().is5xxServerError

    protected fun WebTestClient.ResponseSpec.expectErrorBody(error: Error): WebTestClient.BodyContentSpec = expectBody()
        .jsonPath("$.statusCode").isEqualTo(error.code)
        .jsonPath("$.error").isEqualTo(error.message)

}