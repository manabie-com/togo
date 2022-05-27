# frozen_string_literal: true

class DeviseCreateUsers < ActiveRecord::Migration[7.0]
  def change
    create_table :users, id: :uuid do |t|
      ## Database authenticatable
      t.string :email,              null: false, default: ''
      t.string :encrypted_password, null: false, default: ''

      t.string :name
      t.integer :task_limit, default: 5

      t.timestamps null: false
    end

    add_index :users, :email, unique: true
  end
end
