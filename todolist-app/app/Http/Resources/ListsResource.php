<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class ListsResource extends JsonResource
{
    /**
     * @param Request $request
     * @return array
     */
    public function toArray($request)
    {
        return array(
            'id'=>$this->id,
            'task'=>$this->task,
            'description'=>$this->description,
            'is_complete'=>$this->is_complete
        );
    }

    public function with($request)
    {
        return [
            'author'=>'david.tunglt@gmail.com',
        ];
    }
}
