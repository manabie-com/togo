class RegistrationsController < Devise::RegistrationsController
  include RackSession
  respond_to :json

  def create
    interact = Registration.call(data: params)

    if interact.success?
      super
      @user = User.where(id: current_user.id).first
    else
      render json: { error: interact.error }, status: 422
    end
  end

  private

  def respond_with(resource, _opts = {})
    register_success && return if resource.persisted?

    register_failed
  end

  def register_success
    if setup_account
      render json: { message: 'Success' }, status: :ok
    else
      render json: { error: { message: 'An error has occurred while setting up user.' } }
    end
  end

  def register_failed
    render json: { error: [user: 'Something went wrong.'] }, status: 422
  end

  def setup_account
    @user&.update(task_limit: Random.rand(5...10))
  end

  def decoded_auth_token
    @decoded_auth_token ||= JsonWebToken.decode(http_auth_header)
  end

  def http_auth_header
    return request.headers['Authorization'].split(' ').last if request.headers['Authorization'].present?

    nil
  end
end