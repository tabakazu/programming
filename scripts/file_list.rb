#! /usr/bin/env ruby

ARGV.each do |path|
	puts Dir.glob(path)
end
