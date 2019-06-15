require 'net/http'

def lambda_handler(event:, context:)
  uri = 'https://www.google.co.jp'
  res = Net::HTTP.get_response(URI.parse(uri))
  puts res.code
  res.code
end

