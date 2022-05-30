require 'rails_helper'

RSpec.describe "Api::Todos", type: :request do
  describe "GET /index" do

    context "When user is logged in" do
      before do
        authorize_user

        get '/api/todos', as: :json
      end

      it "returns http success" do
        expect(response.status).to eq(200)
      end
    end

    context "When user is not logged in" do
      before do
        get '/api/todos', as: :json
      end

      it "returns http unauthorized" do
        expect(response.status).to eq(401)
      end
    end

    context "creates todo" do
      let(:user) { authorize_user }

      before do
        params = {
          "todo": {
            "user": user,
            "title": Faker::Lorem.sentence,
            "body": Faker::Lorem.paragraph,
          }
        }

        post '/api/todos', params: params, as: :json
      end

      it "returns http created" do
        expect(response.status).to eq(200)
      end

      it "returns json" do
        hash_body = nil
        expect { hash_body = JSON.parse(response.body).with_indifferent_access }.not_to raise_exception
      end
    end

    context "unauthorized post" do
      before do
        post '/api/todos', params: {}, as: :json
      end

      it "returns http unauthorized" do
        expect(response.status).to eq(401)
      end
    end

  end
end
