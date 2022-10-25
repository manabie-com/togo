<?php
declare(strict_types=1);

namespace App\Service;

use Predis\Client;

/**
 * Class AbstractService
 * @package App\Service
 */
class RedisService
{
    public const PROJECT_NAME = 'todo-test';

    /**
     * @var Client
     */
    private Client $redis;

    /**
     * @param Client $redis
     */
    public function __construct(Client $redis)
    {
        $this->redis = $redis;
    }

    /**
     * @param string $value
     * @return string
     */
    public function generateKey(string $value): string
    {
        return self::PROJECT_NAME . ':' . $value;
    }

    /**
     * @param string $key
     * @return int
     */
    public function exists(string $key): int
    {
        return $this->redis->exists($key);
    }

    /**
     * @param string $key
     * @return object
     */
    public function get(string $key): object
    {
        return json_decode((string) $this->redis->get($key));
    }

    /**
     * @param string $key
     * @param object $value
     */
    public function set(string $key, object $value): void
    {
        $this->redis->set($key, json_encode($value));
    }

    /**
     * @param string $key
     * @param object $value
     * @param int $ttl
     */
    public function setex(string $key, object $value, int $ttl = 3600): void
    {
        $this->redis->setex($key, $ttl, json_encode($value));
    }

    /**
     * @param array $keys
     */
    public function del(array $keys): void
    {
        $this->redis->del($keys);
    }
}
