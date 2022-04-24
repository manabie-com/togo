class IndexUserOnIpAddress < ActiveRecord::Migration[6.0]
  def change
    add_index :users, :remote_ip, unique: true
  end
end
