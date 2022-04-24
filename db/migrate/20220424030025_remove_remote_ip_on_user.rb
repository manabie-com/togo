class RemoveRemoteIpOnUser < ActiveRecord::Migration[6.0]
  def change
    remove_column :users, :remote_ip
  end
end
