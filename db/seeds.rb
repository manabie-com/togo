# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: "Star Wars" }, { name: "Lord of the Rings" }])
#   Character.create(name: "Luke", movie: movies.first)


3.times do
  user = User.new(name: Faker::Games::Dota.hero, task_limit: rand(1..3))

  if user.save
    Task.create(
      title: Faker::Book.title,
      body: Faker::Lorem.sentence,
      user_id: user.id
    )
  end
end