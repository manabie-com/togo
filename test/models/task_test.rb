# frozen_string_literal: true

require 'test_helper'

class TaskTest < ActiveSupport::TestCase
  test 'invalid task' do
    task = Task.new(name: 'Task')
    assert_not task.valid?
    assert_equal task.errors.full_messages, ['User must exist']
  end
end
