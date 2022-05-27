class ApplicationController < ActionController::API
  include ActionController::ImplicitRender

  before_action :configure_permitted_parameters, if: :devise_controller?

  protected
  def configure_permitted_parameters
    devise_parameter_sanitizer.permit(:sign_up, keys: %i[password_confirmation name])
  end
end
