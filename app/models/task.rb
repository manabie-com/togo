# == Schema Information
#
# Table name: tasks
#
#  id         :bigint           not null, primary key
#  name       :string
#  users_id   :bigint           not null
#  created_at :datetime         not null
#  updated_at :datetime         not null
#
class Task < ApplicationRecord
  belongs_to :user
end
