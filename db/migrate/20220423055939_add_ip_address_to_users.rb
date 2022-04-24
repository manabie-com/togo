class AddIpAddressToUsers < ActiveRecord::Migration[6.0]
  def change
    add_column :users, :remote_ip, :string, unique: true
  end
end
