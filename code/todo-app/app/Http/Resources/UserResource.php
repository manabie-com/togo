<?php

namespace App\Http\Resources;

use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class UserResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @param  Request $request
     * @return array
     */
    public function toArray($request): array
    {
        /** @var User $resource */
        $resource = $this->resource;

        if (!$resource instanceof User) {
            return [];
        }

        return [
            'id' => $resource->id,
            'username' => $resource->username,
            'max_todo' => $resource->max_todo,
            'created_at' => $resource->created_at
        ];
    }
}
