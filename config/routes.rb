Rails.application.routes.draw do
  get 'rejection', to: 'rejections#show'
  post 'rejection', to: 'rejections#create'
end
