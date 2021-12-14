require 'rails_helper'

describe "login and register process" do

  it 'should create a user if not exist' do
    visit root_path
    fill_in "user_email",	with: "test@example.com"
    fill_in "user_password",	with: "test@example.com1"
    click_button 'Login/Register'
    expect(page).to have_content 'Logout'
  end
end
