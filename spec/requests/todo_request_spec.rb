require "rails_helper"

RSpec.describe Todo, type: :request do
  describe "POST /api/v1/todos" do
    context "user is unauthenticated and has not reached daily limit" do
      it "creates todo" do
        expect {
          post "/api/v1/todos", :params => {
            :todos => {
              :title => "title",
              :content => "content",
              :done => true
            }
          }

          expect(response).to have_http_status(201)

          data = JSON.parse(response.body)
          expect(data["status"]).to eq 201
          expect(data["message"]).to eq "Todo created"
        }.to change { Todo.count }.by 1
      end
    end

    context "user is unauthenticated and has reached daily limit" do
      it "does not create todo" do
        set_client_to_reach_daily_limit

        expect {
          post "/api/v1/todos", :params => {
            :todos => {
              :title => "title",
              :content => "content",
              :done => true
            }
          },
          env: { "REMOTE_ADDR": "1.2.3.4" }

          expect(response).to have_http_status(429)

          data = JSON.parse(response.body)
          expect(data["status"]).to eq 429
          expect(data["message"]).to eq "Too many requests"
        }.not_to change { Todo.count }
      end

      def set_client_to_reach_daily_limit
        remote_ip = "1.2.3.4"
        date = Time.now.strftime("%d:%m:%Y")
        client_id = "#{remote_ip}:#{date}"
        redis = Redis.new
        redis.set(client_id, 30)
      end
    end
  end
end
