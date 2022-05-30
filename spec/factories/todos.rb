FactoryBot.define do
  factory :todo do
    user { nil }
    title { Faker::Lorem.sentence }
    body { Faker::Lorem.paragraph }
  end
end
