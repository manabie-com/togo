class User < ApplicationRecord
  has_many :todos
  validate :email, unique: true
end
