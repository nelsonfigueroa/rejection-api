class RejectionsController < ApplicationController

  # GET /rejections
  def show
    render json: Rejection.all.sample
  end

  # POST /rejections
  def create
    @rejection = Rejection.new(rejection_params)

    if @rejection.save
      render json: @rejection, status: :created#, location: @rejection
    else
      render json: @rejection.errors, status: :unprocessable_entity
    end
  end

  private
    def rejection_params
      params.permit(:message)
    end
end
