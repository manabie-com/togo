# frozen_string_literal: true
class TasksController < ApplicationController
  before_action :get_user

  def create
    if @user.reach_daily_task_limit?
      render json: { error: 'This user has reached its maximum daily limit of adding tasks!' }
    else
      task = Task.create!(task_params)
      render json: task
    end
  end

  private

  def get_user
    @user = User.find_by(id: task_params[:user_id])

    render json: { error: 'User not found!' } unless @user
  end

  def task_params
    params.permit(:name, :user_id)
  end
end