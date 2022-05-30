# frozen_string_literal: true

module Api
  module Todos
    class UpdateTodoValidator
      include Helper::BasicHelper
      include ActiveModel::API

      attr_accessor(
        :user_id,
        :todo_id,
        :title,
        :body
      )

      validate :user_id_exist, :todo_id_exist, :todo_exist, :required

      def submit
        init
        persist!
      end

      private

      def init
        @user = User.where(id: user_id).load_async.first
        @todo = Todo.where(id: todo_id).load_async.first
      end

      def persist!
        return true if valid?

        false
      end

      def required
        errors.add(:user_id, REQUIRED_MESSAGE) if user_id.blank?
        errors.add(:todo_id, REQUIRED_MESSAGE) if todo_id.blank?
        errors.add(:title, REQUIRED_MESSAGE) if title.blank?
        errors.add(:body, REQUIRED_MESSAGE) if body.blank?
      end

      # Check user if exists.
      def user_id_exist
        errors.add(:user_id, USER_ID_NOT_FOUND) unless @user
      end

      # Check todo if exists.
      def todo_id_exist
        errors.add(:user_id, NOT_FOUND) unless @todo
      end

      # Check record if exists.
      def todo_exist
        todo_exists = Todo.where(title: title, body: body).first
        errors.add(:todo, RECORD_EXIST_MESSAGE) if todo_exists && !todo_exists.id.eql?(@todo.id)
      end
    end
  end
end
