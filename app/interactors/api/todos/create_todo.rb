class Api::Todos::CreateTodo
  include Interactor

  delegate :data, :current_user, to: :context

  def call
    validate!
    build
  end

  def rollback
    context.todo&.destroy
  end

  private

  def build
    @todo = Todo.new(payload)
    Todo.transaction do
      @todo.save
    end

    context.todo = @todo
  end

  def validate!
    verify = Api::Todos::CreateTodoValidator.new(payload)
    return true if verify.submit

    context.fail!(error: verify.errors)
  end

  def payload
    {
      user_id: current_user.id,
      title: data[:todo][:title],
      body: data[:todo][:body]
    }
  end
end
