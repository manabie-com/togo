class LoginController < Devise::SessionsController

  # POST /resource/sign_in
  def create
    if User.where(email: sign_in_params[:email]).first.blank?
      # register new user if not exist
      user = User.new({
        :email => sign_in_params[:email],
        :password => sign_in_params[:password], 
        :password_confirmation => sign_in_params[:password] 
      }).save(:validate => false)
    end

    super
  end
end
