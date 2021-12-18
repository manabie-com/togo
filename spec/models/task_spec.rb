require 'rails_helper'
require 'factories'

RSpec.describe Task, :type => :model do
  it "is valid with valid attributes" do
    user = FactoryBot.create(:user)
    expect(FactoryBot.create(:task, user_id: user.id).save).to be_truthy
  end

  it "is invalid without a user" do
    expect(FactoryBot.build(:task).valid?).to be_falsey
  end

  it "is invalid without a title" do
    user = FactoryBot.create(:user)
    expect(FactoryBot.build(:task, title: nil, user_id: user.id).valid?).to be_falsey
  end

  it "is invalid without a description" do
    user = FactoryBot.create(:user)
    expect(FactoryBot.build(:task, description: nil, user_id: user.id).valid?).to be_falsey
  end

end
