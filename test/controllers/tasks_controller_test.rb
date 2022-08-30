# frozen_string_literal: true

require 'test_helper'

class TasksControllerTest < ActionDispatch::IntegrationTest
  fixtures :users

  test 'can add new task when provide correct params' do
    post '/tasks/create', params: { name: 'Test', user_id: users(:user1).id }
    assert_response 200

    body = JSON.parse(response.body)
    assert_equal body['name'], 'Test'
    assert_equal body['user_id'], users(:user1).id
  end

  test 'return 400 error when provide invalid/missing params' do
    post '/tasks/create', params: { name: 'Test' }
    assert_response 400

    body = JSON.parse(response.body)
    assert_equal body['error'], 'User not found!'

    # Since the fixtures only initialize 2 users, meaning we only have user_id being 1 or 2
    post '/tasks/create', params: { name: 'Test', user_id: 400 }
    assert_response 400

    body = JSON.parse(response.body)
    assert_equal body['error'], 'User not found!'
  end

  test 'return error when user has reached daily task limit' do
    user = users(:user1)
    Task.create(name: 'Task 1', user_id: user.id)
    Task.create(name: 'Task 2', user_id: user.id)
    Task.create(name: 'Task 3', user_id: user.id)

    post '/tasks/create', params: { name: 'Test', user_id: user.id }
    assert_response 200

    body = JSON.parse(response.body)
    assert_equal body['error'], "This user has reached its maximum daily limit of adding tasks! (#{user.max_daily_tasks} tasks/day)"
  end

  test 'user can add new tasks the next day' do
    user = users(:user1)
    Task.create(name: 'Task 1', user_id: user.id)
    Task.create(name: 'Task 2', user_id: user.id)
    Task.create(name: 'Task 3', user_id: user.id)

    post '/tasks/create', params: { name: 'Test', user_id: user.id }
    assert_response 200

    body = JSON.parse(response.body)
    assert_equal body['error'], "This user has reached its maximum daily limit of adding tasks! (#{user.max_daily_tasks} tasks/day)"

    travel_to(Time.current + 1.day) do
      post '/tasks/create', params: { name: 'Next day Task', user_id: user.id }
      assert_response 200

      body = JSON.parse(response.body)
      assert_equal body['name'], 'Next day Task'
    end
  end
end
