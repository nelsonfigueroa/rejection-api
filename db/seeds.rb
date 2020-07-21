require 'csv'
rejections = CSV.read("db/rejections.csv")

rejections.each do |rejection|
  Rejection.create(message: rejection[1])
end

# for some reason there's a nil entry at the end
Rejection.last.destroy