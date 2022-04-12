<?php

namespace Support\Sort;

/*
 * Use sort
 * https://spatie.be/docs/laravel-query-builder/v3/features/sorting
 */

abstract class Sort
{
    public string $defaultSort = '';

    abstract function allowedSorts(): array;
}
