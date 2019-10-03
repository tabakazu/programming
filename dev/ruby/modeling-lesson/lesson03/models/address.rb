class Address
  attr_accessor :id, :customer_id, :name, :street_address

  def initialize(id, customer_id, name, street_address)
    @id = id
    @customer_id = customer_id
    @name = name
    @street_address = street_address
  end
end
