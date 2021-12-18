class CreateTasks < ActiveRecord::Migration[6.1]
  def change
    create_table :tasks do |t|
      t.string :title, null: false
      t.text :description, null: false

      t.belongs_to :user

      t.timestamps
    end
  end
end
