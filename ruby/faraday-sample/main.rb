require 'faraday'

res = Faraday.get 'https://tabakazu.com'
p res.status
# => 200
p res.headers
# => {"content-type"=>"text/html", "content-length"=>"2048", ... }
p res.body
# => "<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"/>...>
