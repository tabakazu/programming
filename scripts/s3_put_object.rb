require 'aws-sdk'

client = Aws::S3::Client.new(
  :region => ENV['AWS_REGION'],
  :access_key_id => ENV['AWS_ACCESS_KEY_ID'],
  :secret_access_key => ENV['AWS_SECRET_ACCESS_KEY']
)

client.put_object(
  :bucket => ENV['AWS_S3_BUCKET'],
  :key => 'memo.txt',
  :body => File.open('memo.txt')
)