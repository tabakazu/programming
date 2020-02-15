require 'minitest/unit'
require 'minitest/autorun'
require './models/customer'

MiniTest::Unit.autorun

class TestCustomer < MiniTest::Unit::TestCase
  def setup
    @customer = Customer.new(1, '鈴木ハナコ', 'hanako@example.jp')
  end

  def test_id
    assert_equal 1, @customer.id
  end

  def test_name
    assert_equal '鈴木ハナコ', @customer.name
  end

  def test_email
    assert_equal 'hanako@example.jp', @customer.email
  end
end
