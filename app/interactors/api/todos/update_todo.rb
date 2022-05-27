class Api::Todos::UpdateTodo
  include Interactor

  delegate :data, :current_user, to: :context

  def call
    init
    validate!
    build
  end

  def rollback; end

  private

  def init
    @todo = Todo.where(id: payload[:todo_id]).first

    context.fail!(error: { todo_id: ['Not found.'] }) unless @todo
    context.fail!(error: { user: ['You do not have access to edit this todo.'] }) unless @todo.user_id.eql?(payload[:user_id])
  end

  def build
    @todo&.update(
      title: payload[:title],
      body: payload[:body]
    )
  end

  def validate!
    verify = Api::Todos::UpdateTodoValidator.new(payload)

    return true if verify.submit

    context.fail!(error: verify.errors)
  end

  def payload
    {
      todo_id: data[:todo_id],
      user_id: current_user.id,
      title: data[:todo][:title],
      body: data[:todo][:body]
    }
  end
end
