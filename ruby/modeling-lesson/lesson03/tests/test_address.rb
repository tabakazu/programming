require 'minitest/unit'
require 'minitest/autorun'
require './models/address'

MiniTest::Unit.autorun

class TestAddress < MiniTest::Unit::TestCase
  def setup
    @address = Address.new(1, 1, '自宅', '中央区銀座1-1')
  end

  def test_id
    assert_equal 1, @address.id
  end

  def test_customer_id
    assert_equal 1, @address.customer_id
  end

  def test_name
    assert_equal '自宅', @address.name
  end

  def test_email
    assert_equal '中央区銀座1-1', @address.street_address
  end
end
