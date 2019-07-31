module Aws
  module Ses
    module Plugin
      autoload :DeliveryMethod, 'aws/ses/plugin/delivery_method'
    end
  end
end

require 'aws/ses/plugin/railtie' if defined?(Rails::Railtie)
