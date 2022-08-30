# frozen_string_literal: true

require 'test_helper'

class UserTest < ActiveSupport::TestCase
  def setup
    @user = User.create!(
      name: 'Test User',
      max_daily_tasks: 3
    )
  end

  test '.reach_daily_task_limit?' do
    # Initially user has not created any task yet
    assert_not @user.reach_daily_task_limit?

    # After user created 3 tasks
    Task.create!(name: 'Task 1', user_id: @user.id)
    Task.create!(name: 'Task 2', user_id: @user.id)
    Task.create!(name: 'Task 3', user_id: @user.id)
    assert @user.reload.reach_daily_task_limit?

    # On the next day, reach_daily_task_limit should reset
    travel_to(Time.current + 1.day) do
      assert_not @user.reach_daily_task_limit?
    end
  end
end
