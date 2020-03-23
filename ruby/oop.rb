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


###
### ポリモーフィズム(多様性)
###
class Customer # 顧客
  # テンプレートパターン
  def info
    raise NotImplementedError, "This #{self.class} cannot respond to:"
  end
end

class IndividualCustomer < Customer # 個人顧客
  def info
    { name: @name, sex: @sex }
  end
end

class CorporateCustomer < Customer # 法人顧客
  def info
    { name: @name, capital: @capital }
  end
end

def get_info_by_customer customer
  customer.info
end

individual_customer_2 = IndividualCustomer.new({ name: 'Bob', sex: 'male' })
corporate_customer_2 = CorporateCustomer.new({ name: 'あいうえお株式会社', capital: 1_000 })

p get_info_by_customer(individual_customer_2)
p get_info_by_customer(corporate_customer_2)
