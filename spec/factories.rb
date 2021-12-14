FactoryBot.define do
  factory :user do
    email { 'test@example.com' }
    password { '42ds4gy' }
  end

  factory :task do
    title { 'The first task' }
  end
end