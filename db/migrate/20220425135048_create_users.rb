class CreateUsers < ActiveRecord::Migration[6.1]
  def change
    create_table :users do |t|
      t.string :name
      t.integer :max_daily_tasks
      
      t.timestamps
    end
  end
end
