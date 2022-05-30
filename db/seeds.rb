# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: "Star Wars" }, { name: "Lord of the Rings" }])
#   Character.create(name: "Luke", movie: movies.first)
user =  User.create(email: 'jsmith@test.com', password: 'Abc!23', name: 'John Smith', task_limit: 5)

Todo.create(title: "Breakfast", body: "Eat Breakfast", user_id: user.id)
Todo.create(title: "Lunch", body: "Eat Lunch", user_id: user.id)
Todo.create(title: "Dinner", body: "Eat Dinner", user_id: user.id)