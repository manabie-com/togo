<?php

namespace Domain\Tasks\Transformers;

use Domain\Tasks\Models\Task;
use Domain\Users\Transformers\UserTransformer;
use Flugg\Responder\Transformers\Transformer;

class TaskTransformer extends Transformer
{
    /**
     * List of available relations.
     *
     * @var string[]
     */
    protected $relations = [
        'user' => UserTransformer::class,
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
     * @param Task $task
     * @return array
     */
    public function transform(Task $task): array
    {
        return [
            'id'          => (string) $task->id,
            'name'        => (string) $task->name,
            'description' => (string) $task->description,
            'user_id'     => (string) $task->user_id,
            'created_at'  => (string) $task->created_at,
        ];
    }
}
