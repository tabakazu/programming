require 'net/http'
require 'uri'

ARGV.each do |url|
  res = Net::HTTP.get_response(URI.parse(url))
  puts res.code
end
