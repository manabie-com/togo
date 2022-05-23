class Task < ApplicationRecord
  belongs_to :user

  validates :title, presence: true
  validates :body, presence: true
  validate :user_limit
  
  def user_limit
    if  user.tasks.where(created_at: Time.zone.now.beginning_of_day..Time.zone.now).length >= user.task_limit
      errors.add(:user_limit, message: "Task limit has been reached for today.")
    end
  end
end
