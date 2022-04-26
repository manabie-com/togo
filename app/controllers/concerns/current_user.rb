module CurrentUser
  extend ActiveSupport::Concern

  def current_user
    User.find_or_create_by(remote_ip: client_remote_ip)
  end
end
