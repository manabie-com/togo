<?php
declare(strict_types=1);

namespace Tests\integration;

/**
 * Class TaskTest
 * @package App\Service
 */
class TaskTest extends BaseTestCase
{
    private static string $id;

    /**
     * Test Get All Tasks.
     * @throws \Throwable
     */
    public function testGetTasks(): void
    {
        $response = $this->runApp('GET', '/api/v1/tasks');

        $result = (string) $response->getBody();

        $this->assertEquals(200, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('success', $result);
        $this->assertStringContainsString('pagination', $result);
        $this->assertStringContainsString('data', $result);
        $this->assertStringNotContainsString('error', $result);
    }

    /**
     * Test Get Tasks By Page.
     * @throws \Throwable
     */
    public function testGetTasksByPage(): void
    {
        $response = $this->runApp('GET', '/api/v1/tasks?page=1&per_page=3');

        $result = (string) $response->getBody();

        $this->assertEquals(200, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('success', $result);
        $this->assertStringContainsString('pagination', $result);
        $this->assertStringContainsString('data', $result);
        $this->assertStringNotContainsString('error', $result);
    }

    /**
     * Test Get Task Not Found.
     */
    public function testGetTaskNotFound(): void
    {
        $response = $this->runApp('GET', '/api/v1/tasks/123456789');

        $result = (string) $response->getBody();

        $this->assertEquals(404, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Create Task.
     * @throws \Throwable
     */
    public function testCreateTask(): void
    {
        $response = $this->runApp(
            'POST',
            '/api/v1/tasks',
            ['content' => 'New Task']
        );

        $result = (string) $response->getBody();

        self::$id = json_decode($result)->message->id;

        $this->assertEquals(201, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('success', $result);
        $this->assertStringContainsString('content', $result);
        $this->assertStringNotContainsString('error', $result);
    }

    /**
     * Test Get Task Created.
     */
    public function testGetTaskCreated(): void
    {
        $response = $this->runApp('GET', '/api/v1/tasks/' . self::$id);

        $result = (string) $response->getBody();

        $this->assertEquals(200, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('success', $result);
        $this->assertStringContainsString('content', $result);
        $this->assertStringNotContainsString('error', $result);
    }

    /**
     * Test Create Task Without id.
     * @throws \Throwable
     */
    public function testCreateTaskWithOutTaskName(): void
    {
        $response = $this->runApp('POST', '/api/v1/tasks');

        $result = (string) $response->getBody();

        $this->assertEquals(400, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Create Task Without Authorization Bearer JWT.
     */
    public function testCreateTaskWithoutBearerJWT(): void
    {
        $auth = self::$jwt;
        self::$jwt = '';
        $response = $this->runApp(
            'POST',
            '/api/v1/tasks',
            ['content' => 'my task']
        );
        self::$jwt = $auth;

        $result = (string) $response->getBody();

        $this->assertEquals(400, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Create Task With Invalid JWT.
     * @throws \Throwable
     */
    public function testCreateTaskWithInvalidJWT(): void
    {
        $auth = self::$jwt;
        self::$jwt = 'invalidToken';
        $response = $this->runApp(
            'POST',
            '/api/v1/tasks',
            ['content' => 'my task']
        );
        self::$jwt = $auth;

        $result = (string) $response->getBody();

        $this->assertEquals(400, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Create Task With Forbidden JWT.
     */
    public function testCreateTaskWithForbiddenJWT(): void
    {
        $auth = self::$jwt;
        self::$jwt = 'Bearer eyJ0eXAiOiJK1NiJ9.eyJzdWIiOiI4Ii';
        $response = $this->runApp(
            'POST',
            '/api/v1/tasks',
            ['content' => 'my task']
        );
        self::$jwt = $auth;

        $result = (string) $response->getBody();

        $this->assertEquals(403, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Update Task.
     */
    public function testUpdateTask(): void
    {
        $response = $this->runApp(
            'PUT',
            '/api/v1/tasks/' . self::$id,
            ['content' => 'Update Task']
        );

        $result = (string) $response->getBody();

        $this->assertEquals(200, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('success', $result);
        $this->assertStringContainsString('content', $result);
        $this->assertStringNotContainsString('error', $result);
    }

    /**
     * Test Update Task Not Found.
     * @throws \Throwable
     */
    public function testUpdateTaskNotFound(): void
    {
        $response = $this->runApp(
            'PUT',
            '/api/v1/tasks/123456789',
            ['content' => 'Task']
        );

        $result = (string) $response->getBody();

        $this->assertEquals(404, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Update Task of Another User.
     */
    public function testUpdateTaskOfAnotherUser(): void
    {
        $response = $this->runApp(
            'PUT',
            '/api/v1/tasks/6',
            ['content' => 'Task']
        );

        $result = (string) $response->getBody();

        $this->assertEquals(404, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }

    /**
     * Test Delete Task.
     * @throws \Throwable
     */
    public function testDeleteTask(): void
    {
        $response = $this->runApp('DELETE', '/api/v1/tasks/' . self::$id);

        $result = (string) $response->getBody();

        $this->assertEquals(204, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('success', $result);
        $this->assertStringNotContainsString('error', $result);
    }

    /**
     * Test Delete Task Not Found.
     * @throws \Throwable
     */
    public function testDeleteTaskNotFound(): void
    {
        $response = $this->runApp('DELETE', '/api/v1/tasks/123456789');

        $result = (string) $response->getBody();

        $this->assertEquals(404, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringContainsString('error', $result);
    }
}
