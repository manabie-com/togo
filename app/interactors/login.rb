class Login
  include Interactor

  delegate :data, to: :context

  def call
    validate!
    build
  end

  private

  def build; end

  def validate!
    verify = LoginValidator.new(payload)

    return true if verify.submit

    context.fail!(error: verify.errors)
  end

  def payload
    {
      email: data[:user][:email],
      password: data[:user][:password]
    }
  end
end
