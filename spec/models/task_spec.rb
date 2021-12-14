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

  it "is invalid without a link" do
    expect(FactoryBot.build(:task, link: nil).valid?).to be_falsey
  end
end
