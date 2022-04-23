module Api
  module V1
    class TodosController < ApplicationController
      def create
        todo = Todo.new(todo_params)

        if todo.save
          render json: { message: 'Todo created', data: todo }, status: :ok
        else
          render json: { message: 'Something went wrong' }, status: :unprocessable_entity
        end
      end

      private

      def todo_params
        params.permit(:title, :content, :done)
      end

      def user
        # todo: current user or find by ip_address
      end
    end
  end
end
