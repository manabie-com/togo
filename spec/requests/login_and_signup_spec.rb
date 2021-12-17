require 'rails_helper'
require 'factories'

describe "login and register process" do
  let(:valid_params) do
    {
      user:{
        email: "test8@example.com",
        password: "sdf"
      }
    }
  end

  let(:wrong_params) do
    {
      user:{
        email: "test8@example.com",
        password: ""
      }
    }
  end

  it 'should fail to login if empty password' do
    post "/api/users/sign_in", params: wrong_params
    expect(response.status).to eq(422)
  end

  it 'should create a user if not exist' do
    post "/api/users/sign_in", params: valid_params
    expect(response.status).to eq(200)
    expect(JSON.parse(response.body)['user']['token']).not_to be_empty
  end
end
