10.times do
  user = User.create({ email: Faker::Internet.email })
  user.todos.create({ title: Faker::Book.title, content: Faker::Lorem.sentence })
end
