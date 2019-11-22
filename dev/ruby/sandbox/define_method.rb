class Greeting
  { 
    hello: %Q(p "Hello"),
    konitiha: %Q(p "こんにちは")
  }.each do |k, v|
    define_method(k) { eval(v) }
  end
end
 
Greeting.new.hello #=> "Hello"
Greeting.new.konitiha #=> "こんにちは"

# ------------------

[:user, :admin].each do |name|
  define_method("current_#{name}?") { true }
end

p current_user? #=> true
p current_admin? #=> true

# ------------------

[:user, :admin].each do |name|
  define_method("#{name}?") { |resource| resource.class.to_s == name.capitalize.to_s }
end

class User; end
p user?(User.new) #=> true

# ------------------

class User
  attr_accessor :role
  def initialize(hash)
    @role = hash[:role]
  end

  [:admin].each do |name|
    define_method("#{name}?") { role == name }
  end
end

admin = User.new(role: :admin)
p admin.admin? #=> true
guest = User.new(role: :guest)
p guest.admin? #=> false
