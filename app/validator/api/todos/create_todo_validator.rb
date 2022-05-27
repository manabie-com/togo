# frozen_string_literal: true

module Api
  module Todos
    class CreateTodoValidator
      include Helper::BasicHelper
      include ActiveModel::API

      attr_accessor(
        :user_id,
        :title,
        :body
      )

      validate :user_id_exist, :todo_exist, :user_limit, :required

      def submit
        init
        persist!
      end

      private

      def init
        @user = User.where(id: user_id).first
      end

      def persist!
        return true if valid?

        false
      end

      def required
        errors.add(:user_id, REQUIRED_MESSAGE) if user_id.blank?
        errors.add(:title, REQUIRED_MESSAGE) if title.blank?
        errors.add(:body, REQUIRED_MESSAGE) if body.blank?
      end

      def user_id_exist
        errors.add(:user_id, USER_ID_NOT_FOUND) unless @user
      end

      def todo_exist
        errors.add(:todo, RECORD_EXIST_MESSAGE) if Todo.exists?(title: title, body: body)
      end

      def user_limit
        user_limit = @user.task_limit ? @user.task_limit : 0
        errors.add(:todo, LIMIT_REACHED) if @user.todos.where(created_at: Time.zone.now.beginning_of_day..Time.zone.now.end_of_day).count >= user_limit
      end
    end
  end
end
