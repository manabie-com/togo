<?php

namespace Api\Requests\Task;

use Api\Requests\ApiRequest;
use Illuminate\Support\Carbon;

class CreateTaskRequest extends ApiRequest
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
        $user = auth()->user();

        return [
            'name'        => 'required|string|max:255',
            'description' => 'nullable|string|max:255',
            'task_number' => 'required|integer|max:'.intval($user->limit_task),
        ];
    }

    protected function prepareForValidation(): void
    {
        $user = auth()->user();
        $currentTask = $user->tasks()->where('created_at', '>=', Carbon::today())->count();
        $this->merge(['task_number' => $currentTask + 1]);
    }
}
