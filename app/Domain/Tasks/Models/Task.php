<?php

namespace Domain\Tasks\Models;

use App\Models\BaseModel;
use Database\Factories\TaskFactory;
use Domain\Tasks\QueryBuilders\TaskQueryBuilder;
use Domain\Users\Models\User;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Support\Traits\HasUuid;

class Task extends BaseModel
{
    use HasUuid;

    protected $fillable = [
        'name',
        'description',
        'user_id',
    ];

    /**
     * Query builder
     *
     * @param $query
     * @return TaskQueryBuilder
     */
    public function newEloquentBuilder($query): TaskQueryBuilder
    {
        $builder = new TaskQueryBuilder($query);
        if ($user = auth()->user()) {
            $builder = $builder->where('user_id', $user->id);
        }

        return $builder;
    }

    /**
     * Factory
     *
     * @return TaskFactory
     */
    protected static function newFactory(): TaskFactory
    {
        return TaskFactory::new();
    }

    /**
     * Owner
     *
     * @return BelongsTo
     */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class);
    }
}
