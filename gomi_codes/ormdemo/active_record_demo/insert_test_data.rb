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

10.times do | x |
  post = Post.create(title: "Post #{x}", content: "Body #{x}")
  2.times do | y |
    post.comments.create(content: "Comment #{y}")
  end
end