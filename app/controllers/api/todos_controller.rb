class Api::TodosController < ApplicationController
  before_action :authenticate_user!
  before_action :set_todo, only: :show

  # GET /api/todos
  def index
    @todos = Todo.where(user_id: current_user.id)
  end

  # GET /api/todos/:id
  def show; end

  # POST /api/todos
  def create
    interact = Api::Todos::CreateTodo.call(data: params, current_user: current_user)

    if interact.success?
      @todo = interact.todo
    else
      render json: { error: interact.error }, status: 422
    end
  end

  # PATCH/PUT /api/todos/:id
  def update
    interact = Api::Todos::UpdateTodo.call(data: params, current_user: current_user)

    if interact.success?
      render json: { message: 'Success' }
    else
      render json: { error: interact.error }, status: 422
    end
  end



  private
  # Get params todo_id and handle if record does not exsits.
  def set_todo
    @todo = Todo.find(params[:todo_id])
  rescue ActiveRecord::RecordNotFound
    render json: { error: { todo_id: ['Not Found.'] } }, status: :not_found
  end
end
