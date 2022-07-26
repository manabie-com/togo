require "rails_helper"

RSpec.describe ::Api::TaskController, type: :controller do

  describe "POST create" do
    context "when user has not reached task limit while creating task" do
      let!(:user){ create :user, name: 'Raffy', task_limit: 5}
      let!(:task_params) do
      {
        title: 'Title1',
        body: 'Body1',
        user_id: user.id
      }
      end

      it "creates a todo task" do
        post :create, params: {
          task: task_params
        }
        expect(response).to have_http_status(:created)
      end
    end

    context "when user has reached task limit while creating task" do
      let!(:user){ create :user, name: 'Raffy', task_limit: 1}
      let!(:task_params) do
      {
        title: 'Title1',
        body: 'Body1',
        user_id: user.id
      }
      end

      before(:example) { create :task, title: 'Title2', body: 'Body2', user_id: user.id}
      it "creates a todo task" do
        post :create, params: {
          task: task_params
        }
        expect(response).to have_http_status(:unprocessable_entity)
      end
    end

    context "when user id not in parameters when creating a task" do
      let!(:task_params) do
      {
        title: 'Title1',
        body: 'Body1'
      }
      end

      it "creates a todo task" do
        post :create, params: {
          task: task_params
        }
        expect(response).to have_http_status(:unprocessable_entity)
      end
    end
  end
end