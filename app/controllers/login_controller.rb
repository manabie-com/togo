class LoginController < Devise::SessionsController
  protect_from_forgery with: :null_session
  before_action :authenticate_user, except: [:new, :create]
  
  # POST /resource/sign_in
  def create
    user = User.find_by_email(sign_in_params[:email])
    unless user
      user = User.new({
        :email => sign_in_params[:email],
        :password => sign_in_params[:password], 
        :password_confirmation => sign_in_params[:password] 
      })
      user.save(:validate => false)
      sign_in(user)
    end

    if user && user.valid_password?(sign_in_params[:password])
      @current_user = user
    else
      render json: {errors: {'email or password' => ['is invalid']}}, status: :unprocessable_entity
    end
  end
end
