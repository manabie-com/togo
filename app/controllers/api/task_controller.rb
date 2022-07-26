class Api::TaskController < ApplicationController
  def index
    tasks = Task.all
    render json: tasks, status: :ok
  end

  def create
    task = Task.create(task_params)

    if task.save
      render json: task, status: :created 
    else
      render json: task.errors, status: :unprocessable_entity
    end
  end

  def task_params
    params.require(:task).permit(:title, :body, :user_id)
  end
end
