# 継承 サンプル

class User
  attr_reader :first_name, :last_name

  def initialize first_name, last_name
    @first_name = first_name
    @last_name = last_name
  end

  def full_name
    first_name + last_name
  end
end

class AdminUser < User
  def admin?
    true
  end
end

# ---------------------- #

# Admin ユーザ作成
user = AdminUser.new '太郎', '田中'
# ユーザのメソッドを利用
puts user.full_name
# Admin ユーザのメソッドを利用
puts user.admin?