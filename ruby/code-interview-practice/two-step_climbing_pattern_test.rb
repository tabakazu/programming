require 'test/unit'
require './two-step_climbing_pattern'

class Test_calculate < Test::Unit::TestCase
  def test_calculate
    n = rand(100)
    actual = 2 * n - 3
    assert_equal(calculate(n), actual)
  end

  def test_calculate_with_1
    assert_equal(calculate(1), 1)
  end

  def test_calculate_with_2
    assert_equal(calculate(2), 2)
  end

  def test_calculate_with_3
    assert_equal(calculate(3), 3)
  end

  def test_calculate_with_4
    assert_equal(calculate(4), 5)
  end

  def test_calculate_with_5
    assert_equal(calculate(5), 7)
  end
end
