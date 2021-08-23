<?php
declare(strict_types=1);

namespace App\Service;

/**
 * Class AbstractService
 * @package App\Service
 */
abstract class AbstractService
{
    protected const DEFAULT_PER_PAGE_PAGINATION = 5;

    protected static function isRedisEnabled(): bool
    {
        return filter_var($_SERVER['REDIS_ENABLED'], FILTER_VALIDATE_BOOLEAN);
    }
}
