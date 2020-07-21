class RejectionsController < ApplicationController

  # GET /rejection
  def show
    render json: Rejection.all.sample.message
  end

  # POST /rejection
  def create
    @rejection = Rejection.new(rejection_params)

    if @rejection.save
      render json: @rejection, status: :created, location: @rejection
    else
      render json: @rejection.errors, status: :unprocessable_entity
    end
  end

  private
    def rejection_params
      params.require(:rejection).permit(:message)
    end
end
