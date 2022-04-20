<?php

namespace Domain\Tasks\Filters;

use Support\Filter\Filter;

class TaskFilter extends Filter
{
    protected array $allowedFilters = [
        'task',
        'description',
    ];
}
