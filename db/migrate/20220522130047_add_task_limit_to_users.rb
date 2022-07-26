class AddTaskLimitToUsers < ActiveRecord::Migration[7.0]
  def change
    add_column :users, :task_limit, :integer
  end
end
