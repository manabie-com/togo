class User < ApplicationRecord
  has_many :todos
  validate :remote_ip, unique: true
end
