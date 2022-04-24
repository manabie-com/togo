require "rails_helper"

RSpec.describe Todo, type: :request do
  include ActiveSupport::Testing::TimeHelpers

  REMOTE_IP = "1.2.3.4"

  describe "POST /api/v1/todos" do
    context "user is unauthenticated and has not reached daily limit" do
      it "creates todo" do
        clear_redis

        expect {
          post "/api/v1/todos", :params => {
            :todos => {
              :title => "title",
              :content => "content",
              :done => true
            }
          }, env: { "REMOTE_ADDR": REMOTE_IP }

          expect(response).to have_http_status(201)

          data = JSON.parse(response.body)
          expect(data["status"]).to eq 201
          expect(data["message"]).to eq "Todo created"
        }.to change { Todo.count }.by 1
      end
    end

    context "user is unauthenticated and has reached daily limit" do
      it "does not create todo" do
        clear_redis
        client_reach_daily_limit

        expect {
          post "/api/v1/todos", :params => {
            :todos => {
              :title => "title",
              :content => "content",
              :done => true
            }
          }, env: { "REMOTE_ADDR": REMOTE_IP }

          expect(response).to have_http_status(429)

          data = JSON.parse(response.body)
          expect(data["status"]).to eq 429
          expect(data["message"]).to eq "Too many requests"
        }.not_to change { Todo.count }
      end
    end

    context "user is unauthenticated and daily limit is reset" do
      it "creates todo" do
        clear_redis
        client_reach_daily_limit

        travel_to 2.days.from_now do
          expect {
            post "/api/v1/todos", :params => {
              :todos => {
                :title => "title",
                :content => "content",
                :done => true
              }
            }, env: { "REMOTE_ADDR": REMOTE_IP }

            expect(response).to have_http_status(201)

            data = JSON.parse(response.body)
            expect(data["status"]).to eq 201
            expect(data["message"]).to eq "Todo created"
          }.to change { Todo.count }.by 1
        end
      end
    end
  end

  def client_reach_daily_limit
    redis = Redis.new
    redis.set(client_id, ENV["unauthenticated_post_todo_limit"].to_i)
  end

  def clear_redis
    redis = Redis.new
    redis.del(client_id)
  end

  def client_id
    remote_ip = REMOTE_IP
    date = Time.now.strftime("%d:%m:%Y")

    "#{remote_ip}:#{date}"
  end
end
