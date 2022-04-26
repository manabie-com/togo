class User < ApplicationRecord
  has_many :todos, dependent: :destroy
end
