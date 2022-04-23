class User < ApplicationRecord
  has_many :todos
  validate :email, unique: true
  validate :remote_ip, unique: true # todo: remove this
end
