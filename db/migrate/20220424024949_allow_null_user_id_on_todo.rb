class AllowNullUserIdOnTodo < ActiveRecord::Migration[6.0]
  def change
    change_column :todos, :user_id, :integer, null: true
  end
end
