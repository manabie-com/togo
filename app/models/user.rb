# == Schema Information
#
# Table name: users
#
#  id              :bigint           not null, primary key
#  name            :string
#  max_daily_tasks :integer
#  created_at      :datetime         not null
#  updated_at      :datetime         not null
#
class User < ApplicationRecord
  has_many :tasks, dependent: :delete_all

  def reach_daily_task_limit?
    tasks.size >= max_daily_tasks
  end
end
