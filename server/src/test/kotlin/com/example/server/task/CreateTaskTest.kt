package com.example.server.task

import com.example.server.BaseTest
import com.example.server.config.database.DatabaseConfigTest
import com.example.server.rest.request.TaskRequest
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.test.context.ContextConfiguration
import org.springframework.test.web.reactive.server.WebTestClient
import org.springframework.web.reactive.function.BodyInserters

@ContextConfiguration(classes = [DatabaseConfigTest::class])
class CreateTaskTest : BaseTest() {
    private val apiPath = "/api/tasks"

    @BeforeEach
    fun setupTest() {
        server.start(1234)
    }

    @AfterEach
    fun tearDown() {
        server.shutdown()
    }

    private fun <T> assertTestResult(
        expectedStatus: HttpStatus,
        body: T,

    ): WebTestClient.BodyContentSpec {
        return clientUseKotlinSerialization.post()
            .uri { uriBuilder ->
                uriBuilder.path(apiPath)
                    .build()
            }

            .accept(MediaType.APPLICATION_JSON)
            .body(BodyInserters.fromValue(body))
            .exchange()
            .expectStatus().let {
                when (expectedStatus) {
                    HttpStatus.BAD_REQUEST -> it.isBadRequest
                    HttpStatus.INTERNAL_SERVER_ERROR -> it.is5xxServerError
                    else -> it.isOk
                }
            }
            .expectHeader()
            .contentType(MediaType.APPLICATION_JSON)
            .expectBody()
    }

    @Test
    fun `createTaks-0 add task success`() {
        val request = TaskRequest(
            title = "test",
            description = "aa",
            userId = "0xaa91b0c2bc854c619253accc3cfc9270"
        )
        assertTestResult(
            expectedStatus = HttpStatus.OK,
            body = request
        )
            .jsonPath("$.data.title").isEqualTo("test")
            .jsonPath("$.data.description").isEqualTo("aa")
            .jsonPath("$.data.userId").isEqualTo("0xaa91b0c2bc854c619253accc3cfc9270")
    }

    @Test
    fun `createTaks-400_001 create task out of limit task`() {
        val request = TaskRequest(
            title = "test",
            description = "aa",
            userId = "0xaa91b0c2bf854c619253accc3cfc9270"
        )
        assertTestResult(
            expectedStatus = HttpStatus.OK,
            body = request
        )
            .jsonPath("$.data.title").isEqualTo("test")
            .jsonPath("$.data.description").isEqualTo("aa")
            .jsonPath("$.data.userId").isEqualTo("0xaa91b0c2bf854c619253accc3cfc9270")

        assertTestResult(
            expectedStatus = HttpStatus.BAD_REQUEST,
            body = request
        )
            .jsonPath("$.statusCode").isEqualTo(400001)

    }

    @Test
    fun `createTaks-400_002 when user is not exist`() {
        val request = TaskRequest(
            title = "test",
            description = "aa",
            userId = "aaa"
        )
        assertTestResult(
            expectedStatus = HttpStatus.BAD_REQUEST,
            body = request
        )
            .jsonPath("$.statusCode").isEqualTo(400002)
    }


}