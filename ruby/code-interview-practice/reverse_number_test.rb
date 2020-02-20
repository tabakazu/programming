require 'test/unit'
require './reverse_number'

class Test_f < Test::Unit::TestCase
  def test_f
    n = rand(1000)
    actual = n.to_s.reverse.to_i
    assert_equal(f(n), actual)
  end

  def test_f_with_1234
    assert_equal(f(1234), 4321)
  end

  def test_f_with_1234567890
    assert_equal(f(1234567890), 987654321)
  end

  def test_f_with_1111
    assert_equal(f(1111), 1111)
  end
end
