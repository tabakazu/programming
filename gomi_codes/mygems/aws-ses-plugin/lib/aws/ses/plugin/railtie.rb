module Aws
  module Ses
    module Plugin
      class Railtie < ::Rails::Railtie
        initializer 'aws-ses-plugin.add_delivery_method' do
          ActiveSupport.on_load :action_mailer do
            ActionMailer::Base.add_delivery_method(
              :aws_ses_plugin,
              Aws::Ses::Plugin::DeliveryMethod
            )
          end
        end
      end
    end
  end
end
