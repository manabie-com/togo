Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  scope :api, defaults: {format: :json} do
    devise_for :users, controllers: {sessions: 'login'}
    resources :tasks
  end
end
