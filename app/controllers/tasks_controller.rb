class TasksController < ApplicationController
  before_action :set_task, only: %i[ show edit update destroy ]

  # GET /tasks or /tasks.json
  def index
    tasks = Task.where(user_id: current_user.id)
    render json: tasks
  end

  # GET /tasks/1 or /tasks/1.json
  def show
    render json: @task
  end

  # POST /tasks or /tasks.json
  def create
    @task = Task.new(task_params)
    @task.user_id = current_user.id
    # @entry = Task.where(created_at: Time.now.beginning_of_day.utc..Time.now.end_of_day.utc).first_or_create!
    # @entry.update_attributes(entry_params)
    if @task.save
      render json: @task
    else
      render json: @task.errors, status: :unprocessable_entity
    end
  end

  # PATCH/PUT /tasks/1 or /tasks/1.json
  def update
      if @task.update(task_params)
        render json: @task
      else
        render json: @task.error, status: :unprocessable_entity
    end
  end

  # DELETE /tasks/1 or /tasks/1.json
  def destroy
    @task.destroy
    render json: {message: "Successfully removed"}
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_task
      @task = Task.where(id: params[:id], user_id: current_user.id).take
    end

    # Only allow a list of trusted parameters through.
    def task_params
      params.require(:task).permit(:title, :description)
    end
end
