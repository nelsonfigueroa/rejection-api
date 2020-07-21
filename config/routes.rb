Rails.application.routes.draw do
  root to: 'rejections#index'
  get '/rejections', to: 'rejections#show'
  post '/rejections', to: 'rejections#create'
end
