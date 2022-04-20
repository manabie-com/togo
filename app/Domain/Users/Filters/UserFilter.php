<?php

namespace Domain\Users\Filters;

use Illuminate\Support\Str;
use Support\Builder\Builder;
use Support\Filter\Filter;

class UserFilter extends Filter
{
    protected array $allowedFilters = [
        'name',
        'user_name',
    ];
}
