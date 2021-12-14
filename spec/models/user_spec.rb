require 'rails_helper'
require 'factories'

RSpec.describe User, :type => :model do
  it "is valid with valid attributes" do
    expect(FactoryBot.build(:user).valid?).to be_truthy
  end

  it "is invalid without an email" do
    expect(FactoryBot.build(:user, email: nil).valid?).to be_falsey
  end

  it "is invalid without a password" do
    expect(FactoryBot.build(:user, password: nil).valid?).to be_falsey
  end
end
