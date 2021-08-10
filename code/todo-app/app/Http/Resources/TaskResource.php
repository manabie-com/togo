<?php

namespace App\Http\Resources;

use App\Models\Task;
use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class TaskResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @param  Request $request
     * @return array
     */
    public function toArray($request): array
    {
        /** @var Task $resource */
        $resource = $this->resource;

        if (!$resource instanceof Task) {
            return [];
        }

        return [
            'id' => $resource->id,
            'user_id' => $resource->user_id,
            'content' => $resource->content,
            'created_at' => $resource->created_at
        ];
    }
}
