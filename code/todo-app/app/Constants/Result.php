<?php

namespace App\Constants;

class Result
{
    const CODE_OK = 0;
    const CODE_FAIL = 1;
    const CODE_TASK_REACH_LIMIT_DAY = 100;

    const MESSAGE = [
        self::CODE_OK => 'success',
        self::CODE_FAIL => 'fail',
        self::CODE_TASK_REACH_LIMIT_DAY => 'You have created enough limit for today'
    ];

    public static function getMessage(int $code): string
    {
        return self::MESSAGE[$code] ?? self::MESSAGE[self::CODE_FAIL];
    }
}
