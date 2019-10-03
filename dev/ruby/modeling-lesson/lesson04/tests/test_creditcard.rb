require 'minitest/unit'
require 'minitest/autorun'
require './models/creditcard'

MiniTest::Unit.autorun

class TestCreditcard < MiniTest::Unit::TestCase
  def setup
    @creditcard = Creditcard.new(1, 1, '424242424242', '2019-07-23', 'HANAKO SUZUKI')
  end

  def test_id
    assert_equal 1, @creditcard.id
  end

  def test_customer_id
    assert_equal 1, @creditcard.customer_id
  end

  def test_card_number
    assert_equal '424242424242', @creditcard.card_number
  end

  def test_expiration
    assert_equal '2019-07-23', @creditcard.expiration
  end

  def test_nominee
    assert_equal 'HANAKO SUZUKI', @creditcard.nominee
  end
end
