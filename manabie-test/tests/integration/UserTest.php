<?php
declare(strict_types=1);

namespace Tests\integration;

/**
 * Class UserTest
 * @package App\Service
 */
class UserTest extends BaseTestCase
{
    /**
     * Test that user login endpoint it is working fine.
     * @throws \Throwable
     */
    public function testLoginUser(): void
    {
        $response = $this->runApp('POST', '/login', ['id' => 'kiennt', 'password' => 'abc123']);

        $result = (string) $response->getBody();

        $this->assertEquals(200, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('status', $result);
        $this->assertStringContainsString('success', $result);
        $this->assertStringContainsString('message', $result);
        $this->assertStringContainsString('Authorization', $result);
        $this->assertStringContainsString('Bearer', $result);
        $this->assertStringContainsString('ey', $result);
        $this->assertStringNotContainsString('error', $result);
        $this->assertStringNotContainsString('Failed', $result);
    }

    /**
     * Test login endpoint with invalid credentials.
     */
    public function testLoginUserFailed(): void
    {
        $response = $this->runApp('POST', '/login', ['id' => 'not exists', 'password' => 'p']);

        $result = (string) $response->getBody();

        $this->assertEquals(400, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('Login failed', $result);
        $this->assertStringContainsString('error', $result);
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringNotContainsString('Authorization', $result);
        $this->assertStringNotContainsString('Bearer', $result);
    }

    /**
     * Test login endpoint without send required field password.
     * @throws \Throwable
     */
    public function testLoginWithoutPasswordField(): void
    {
        $response = $this->runApp('POST', '/login', ['id' => 'kiennt']);

        $result = (string) $response->getBody();

        $this->assertEquals(400, $response->getStatusCode());
        $this->assertEquals('application/problem+json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('error', $result);
        $this->assertStringNotContainsString('success', $result);
        $this->assertStringNotContainsString('Authorization', $result);
        $this->assertStringNotContainsString('Bearer', $result);
    }
}
