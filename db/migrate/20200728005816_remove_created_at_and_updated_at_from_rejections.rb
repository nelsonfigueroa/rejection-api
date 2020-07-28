class RemoveCreatedAtAndUpdatedAtFromRejections < ActiveRecord::Migration[6.0]
  def change
    remove_column :rejections, :created_at
    remove_column :rejections, :updated_at
  end
end
