require 'minitest/unit'
require 'minitest/autorun'
require './models/customer'

MiniTest::Unit.autorun

class TestCustomer < MiniTest::Unit::TestCase
  def setup
    @customer = Customer.new(1)
  end

  def test_id
    assert_equal 1, @customer.id
  end
end
