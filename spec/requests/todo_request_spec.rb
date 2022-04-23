require 'rails_helper'

RSpec.describe Todo, type: :request do
  describe "POST /todo" do
    it "done" do
      expect {
        post "/api/v1/todos", :params => {
          :todos => {
            :title => "title",
            :content => "content"
          }
        }

        expect(response).to have_http_status(:created)
      }.to change { Todo.count }.by 1
    end
  end
end
