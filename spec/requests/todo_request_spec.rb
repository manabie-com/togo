require "rails_helper"

RSpec.describe Todo, type: :request do
  include ActiveSupport::Testing::TimeHelpers

  REMOTE_IP = "1.2.3.4"

  describe "POST /api/v1/todos" do
    it "creates new user for a new ip address" do
      clear_redis

      expect(User.where(remote_ip: REMOTE_IP)).to be_empty

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

      expect(User.where(remote_ip: REMOTE_IP)).not_to be_empty

      clear_redis
    end

    context "user has not reached daily limit" do
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

    context "user has reached daily limit" do
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

    context "user daily limit is reset" do
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
    redis.set(client_id, client_post_request_daily_limit)
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

  def client_post_request_daily_limit
    ENV["base_post_request_daily_limit"].to_i + current_user.id
  end

  def current_user
    User.find_or_create_by(remote_ip: REMOTE_IP)
  end
end
