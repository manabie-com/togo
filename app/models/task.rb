class Task < ApplicationRecord
  validates :title, :description, presence: true
  validate :limit_by_user_per_day
  belongs_to :user

  private
  def limit_by_user_per_day
    @entries = Task.where(
      user_id: self.user_id,
      created_at: Time.now.beginning_of_day.utc..Time.now.end_of_day.utc
    )
    if @entries.count > ENV['MAX_TASK_CREATE_BY_USER_PER_DAY'].to_i
      errors.add(:user_id, "User exceeds the limit of tasks create per day")
    end
  end
end
