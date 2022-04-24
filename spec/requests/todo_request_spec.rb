require "rails_helper"

RSpec.describe Todo, type: :request do
  describe "POST /api/v1/todos" do
    context "user is unauthenticated" do
      it "creates todo" do
        expect {
          post "/api/v1/todos", :params => {
            :todos => {
              :title => "title",
              :content => "content",
              :done => true
            }
          }

          expect(response).to have_http_status(:created)

          data = JSON.parse(response.body)
          expect(data["status"]).to eq 201
          expect(data["message"]).to eq "Todo created"
        }.to change { Todo.count }.by 1
      end
    end
  end
end
