class CreateRejections < ActiveRecord::Migration[6.0]
  def change
    create_table :rejections do |t|
      t.text :message

      t.timestamps
    end
  end
end
