class Creditcard
  attr_accessor :id, :customer_id, :card_number, :expiration, :nominee

  def initialize(id, customer_id, card_number, expiration, nominee)
    @id = id
    @customer_id = customer_id
    @card_number = card_number
    @expiration = expiration
    @nominee = nominee
  end
end
