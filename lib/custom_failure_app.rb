class CustomFailureApp < Devise::FailureApp
  def respond
    self.status = 401
    self.content_type = 'application/json'
    self.response_body = { error: { user: ['Access denied!. Token has expired.'] } }.to_json
  end
end