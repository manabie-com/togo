module Api
  module V1
    class TodosController < ApplicationController
      def create
        todo = if user
                 user.todos.new(todo_params)
               else
                 Todo.new(todo_params)
               end

        if todo.save
          render json: { status: 201, message: "Todo created", data: todo },
            status: :created
        else
          render json: { status: 422, message: "Something went wrong" },
            status: :unprocessable_entity
        end
      end

      private

      def todo_params
        params.permit(:title, :content, :done)
      end

      def user
        # current_user
      end

      def client_remote_ip
        request.remote_ip
      end
    end
  end
end
