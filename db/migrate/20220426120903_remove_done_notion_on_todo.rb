class RemoveDoneNotionOnTodo < ActiveRecord::Migration[6.0]
  def change
    remove_column :todos, :done
    remove_column :todos, :done_at
  end
end
