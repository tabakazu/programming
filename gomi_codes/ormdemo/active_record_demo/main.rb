require 'active_record'
require 'erb'

config = YAML::load(ERB.new(IO.read(File.join(File.expand_path('../config', __FILE__), 'database.yml'))).result)
ActiveRecord::Base.establish_connection(config['development'])

class Post < ActiveRecord::Base
  has_many :comments
end

class Comment < ActiveRecord::Base
  belongs_to :post
end

p Post.count
p Comment.count