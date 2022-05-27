class CreateTodos < ActiveRecord::Migration[7.0]
  def change
    create_table :todos, id: :uuid do |t|
      t.references :user, null: false, foreign_key: true, type: :uuid
      t.string :title
      t.string :body

      t.timestamps
    end
  end
end
