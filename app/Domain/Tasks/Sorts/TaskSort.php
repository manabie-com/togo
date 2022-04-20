<?php

namespace Domain\Tasks\Sorts;

use Support\Sort\Sort;

class TaskSort extends Sort
{
    public string $defaultSort = '-updated_at';

    function allowedSorts(): array
    {
        return [
            'task',
            'updated_at',
        ];
    }
}
