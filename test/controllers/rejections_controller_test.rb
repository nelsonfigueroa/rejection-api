require 'test_helper'

class RejectionsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @rejection = rejections(:one)
  end

  test "should get index" do
    get rejections_url, as: :json
    assert_response :success
  end

  test "should create rejection" do
    assert_difference('Rejection.count') do
      post rejections_url, params: { rejection: { message: @rejection.message } }, as: :json
    end

    assert_response 201
  end

  test "should show rejection" do
    get rejection_url(@rejection), as: :json
    assert_response :success
  end

  test "should update rejection" do
    patch rejection_url(@rejection), params: { rejection: { message: @rejection.message } }, as: :json
    assert_response 200
  end

  test "should destroy rejection" do
    assert_difference('Rejection.count', -1) do
      delete rejection_url(@rejection), as: :json
    end

    assert_response 204
  end
end
