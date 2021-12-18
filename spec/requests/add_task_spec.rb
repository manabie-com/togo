require 'rails_helper'
require 'factories'

describe "create task process" do
    let(:task_params) do{
      task: {
        title: "Test title",
        description: "Example description"
      }
    }
    end
    let(:user_params) do
      {
        user:{
          email: "test8@example.com",
          password: "123"
        }
      }
    end

  it 'should fail to create a task without token' do
    post '/api/tasks', params: task_params
    expect(response.status).to eq(401)
  end

  it 'should create a task with token' do
    post "/api/users/sign_in", params: user_params
    post '/api/tasks',
     params: task_params, 
     headers: {:Authorization => "Bearer " + JSON.parse(response.body)['user']['token']}

    expect(response.status).to eq(200)
    expect(JSON.parse(response.body)['title']).to eq('Test title')
  end

  it 'should fail create more than 5 tasks' do
    post "/api/users/sign_in", params: user_params
    token = JSON.parse(response.body)['user']['token']
    10.times do |i|
      post '/api/tasks',
        params: task_params, 
        headers: {:Authorization => "Bearer " + token}
      if i < 4
        expect(response.status).to eq(200)
        expect(JSON.parse(response.body)['title']).to eq('Test title')
      else
        expect(response.status).to eq(422)
      end
    end
  end

end
