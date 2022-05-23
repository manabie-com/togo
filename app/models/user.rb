class User < ApplicationRecord
  has_many :tasks

  validates :task_limit, presence: true
end
