#! /usr/bin/env ruby

require 'net/http'
require 'uri'
require 'date'

ARGV.each do |url|
	begin
	  res = Net::HTTP.get_response(URI.parse(url))
		puts "HTTP/#{res.http_version} #{res.code} #{res.msg}"
		puts "Date: #{Time.now.strftime('%Y/%m/%d %H:%M:%S')}"
	  puts ""
	rescue => e
		puts e.message
	end
end
