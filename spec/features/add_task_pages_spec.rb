require 'rails_helper'
require 'factories'

describe "create task process" do
  before :each do
    user = FactoryBot.create(:user)
    login_as(user)
  end

  it 'should create a task' do
    visit new_task_path
    fill_in "Title",	with: "The first task"
    click_button 'Submit'
    expect(page).to have_content 'Task was successfully created'
  end

  it 'should fail to create a task' do
    visit new_task_path
    click_button 'Submit'
    expect(page).to have_content "Title can't be blank"
  end
end
