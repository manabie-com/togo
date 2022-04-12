<?php

namespace Api\Requests\Task;

use Api\Requests\ApiRequest;

class UpdateTaskRequest extends ApiRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules(): array
    {
        return [
            'name'        => 'sometimes|required|string|max:255',
            'description' => 'sometimes|nullable|string|max:255',
        ];
    }
}
