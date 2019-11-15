require 'active_support/all'

sin = 'person'
mul = 'people'

p sin.pluralize #=> "people"
p mul.singularize #= "person"

# # Gemfile
# gem 'activesupport'
#
# # 実行
# $ bundle
# $ bundle exec ruby sin2mul.rb # or $ ruby sin2mul.rb
