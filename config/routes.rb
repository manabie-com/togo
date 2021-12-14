Rails.application.routes.draw do
  resources :tasks
  root to: 'home#index'
  devise_for :users, controllers: { sessions: 'login'}
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
