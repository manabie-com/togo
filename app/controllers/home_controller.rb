class HomeController < ApplicationController
  skip_before_action :authenticate_user!

  def index
    @tasks = Tasks.all.order(created_at: :desc)
  end
end
