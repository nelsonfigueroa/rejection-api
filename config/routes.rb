Rails.application.routes.draw do
  get '/rejections', to: 'rejections#show'
  post '/rejections', to: 'rejections#create'
end
