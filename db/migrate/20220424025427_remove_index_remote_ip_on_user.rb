class RemoveIndexRemoteIpOnUser < ActiveRecord::Migration[6.0]
  def change
    remove_index :users, name: "index_users_on_remote_ip"
  end
end
