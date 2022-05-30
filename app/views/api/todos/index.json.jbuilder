json.data do
  json.array! @todos.each do |data|
    json.id data.id
    json.title data.title
    json.body data.body
  end
end