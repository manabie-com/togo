module Api
  module V1
    class TodosController < ApplicationController
      include PostRequestTrackable

      before_action :count_request, only: :create

      def create
        if client_reach_daily_limit?
          render json: { status: 429, message: "Too many requests" },
            status: 429
          return
        end

        todo = current_user.todos.new(todo_params)

        if todo.save
          render json: { status: 201, message: "Todo created", data: todo },
            status: 201
        else
          render json: { status: 422, message: "Something went wrong" },
            status: 422
        end
      end

      private

      def todo_params
        params.permit(:title, :content)
      end
    end
  end
end
