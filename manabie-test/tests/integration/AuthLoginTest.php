<?php
declare(strict_types=1);

namespace Tests\integration;

/**
 * Class AuthLoginTest
 * @package App\Service
 */
class AuthLoginTest extends BaseTestCase
{
    /**
     * Test user login endpoint and get a JWT Bearer Authorization.
     * @throws \Throwable
     */
    public function testLogin(): void
    {
        $response = $this->runApp('POST', '/login', ['id' => 'kiennt', 'password' => 'abc123']);

        $result = (string) $response->getBody();

        self::$jwt = json_decode($result)->message->Authorization;

        $this->assertEquals(200, $response->getStatusCode());
        $this->assertEquals('application/json', $response->getHeaderLine('Content-Type'));
        $this->assertStringContainsString('Authorization', $result);
        $this->assertStringContainsString('Bearer', $result);
    }
}
