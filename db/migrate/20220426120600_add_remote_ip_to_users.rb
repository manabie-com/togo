class AddRemoteIpToUsers < ActiveRecord::Migration[6.0]
  def change
    add_column :users, :remote_ip, :string
    add_index :users, :remote_ip, unique: true
  end
end
