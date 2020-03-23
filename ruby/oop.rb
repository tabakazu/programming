###
### 継承
###
class Customer # 顧客
  def initialize attribute
    @name = attribute[:name]
    customer_initialize attribute
  end

  def customer_initialize attribute # フックメッセージ
  end
end

class IndividualCustomer < Customer # 個人顧客
  def customer_initialize attribute
    @sex = attribute[:sex]
  end
end

class CorporateCustomer < Customer # 法人顧客
  def customer_initialize attribute
    @capital = attribute[:capital]
  end
end

individual_customer_1 = IndividualCustomer.new({ name: 'Alice', sex: 'female' })
corporate_customer_1 = CorporateCustomer.new({ name: '株式会社ABC', capital: 1_000_000 })

p individual_customer_1
p corporate_customer_1
