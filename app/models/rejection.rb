class Rejection < ApplicationRecord
  validates :message, presence: true
end
