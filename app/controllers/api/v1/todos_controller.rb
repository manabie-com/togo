module Api
  module V1
    class TodosController < ApplicationController
      def create
        todo = user.todos.new(todo_params)

        if todo.save
          render json: { message: 'Todo created', data: todo }, status: :created
        else
          render json: { message: 'Something went wrong' }, status: :unprocessable_entity
        end
      end

      private

      def todo_params
        params.permit(:title, :content, :done)
      end

      def user
        # todo: auth user
        @user ||= User.find_or_create_by(remote_ip: client_remote_ip)
      end

      def client_remote_ip
        request.remote_ip
      end
    end
  end
end
