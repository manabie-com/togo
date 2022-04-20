<?php

namespace Domain\Users\Transformers;

use Domain\Tasks\Transformers\TaskTransformer;
use Domain\Users\Models\User;
use Flugg\Responder\Transformers\Transformer;

class UserTransformer extends Transformer
{
    /**
     * List of available relations.
     *
     * @var string[]
     */
    protected $relations = [
        'tasks' => TaskTransformer::class,
    ];

    /**
     * List of autoloaded default relations.
     *
     * @var array
     */
    protected $load = [];

    /**
     * Transform the model.
     *
     * @param User $user
     * @return array
     */
    public function transform(User $user): array
    {
        return [
            'id'         => (string) $user->id,
            'user_name'  => (string) $user->user_name,
            'name'       => (string) $user->name,
            'limit_task' => (int) $user->limit_task,
            'created_at' => (string) $user->created_at,
        ];
    }
}
