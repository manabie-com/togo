<?php
declare(strict_types=1);

namespace App\Service\User;

use App\Repository\UserRepository;
use App\Service\AbstractService;
use App\Service\RedisService;

/**
 * Class AbstractUserService
 * @package App\Service\User
 */
abstract class AbstractUserService extends AbstractService
{
    /**
     * @var UserRepository
     */
    protected UserRepository $userRepository;

    /**
     * @var RedisService
     */
    protected RedisService $redisService;

    /**
     * @param UserRepository $userRepository
     * @param RedisService $redisService
     */
    public function __construct(UserRepository $userRepository, RedisService $redisService)
    {
        $this->userRepository = $userRepository;
        $this->redisService = $redisService;
    }
}
