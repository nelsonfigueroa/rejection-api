require 'csv'
rejections = CSV.read('db/rejections.csv')

rejections.each do |rejection|
  Rejection.create(message: rejection[1])
end