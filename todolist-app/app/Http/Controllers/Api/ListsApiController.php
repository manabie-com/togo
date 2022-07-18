<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\ListsResource;
use App\Models\TodoList;
use Illuminate\Http\Request;

class ListsApiController extends Controller
{
    /**
     * @param Request $request
     * @param TodoList $lists
     * @return ListsResource|void
     */
    public function store(Request $request, TodoList $lists)
    {
        //$request->expectsJson($request->all());
        $lists = $request->isMethod('PUT') ? $lists->findOrFail($request->id) : new TodoList();
        $lists->task = $request->task;
        $lists->description = $request->description;
        $lists->is_complete = $request->is_complete;
        if($lists->save())
        {
            return new ListsResource($lists);
        }
    }
}
