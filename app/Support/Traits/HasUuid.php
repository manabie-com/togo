<?php

namespace Support\Traits;

use Illuminate\Support\Str;

trait HasUuid
{
    /**
     * Boot function from laravel.
     */
    protected static function bootHasUuid()
    {
        static::creating(function ($model) {
            $model->keyType = 'string';
            $model->incrementing = false;
            $model->{$model->getKeyName()} = $model->{$model->getKeyName()} ?: (string) Str::orderedUuid();
        });
    }

    /**
     * @inheritDoc
     */
    public function getIncrementing()
    {
        return false;
    }

    /**
     * @inheritDoc
     */
    public function getKeyType()
    {
        return 'string';
    }
}
