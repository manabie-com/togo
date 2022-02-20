<?php


namespace App\Http\Services;


class TodoCounterService
{
    public function checkUpdatedToday($todoCounter) {
        return $todoCounter->updated_at->isToday();
    }

    public function hasReachedMaximum($todoCounter) {
        return $todoCounter->count_today >= $todoCounter->max_count;
    }
}
