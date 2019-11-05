require 'webrick'

server = WEBrick::HTTPServer.new(
  DocumentRoot: './',
  BindAddress: '127.0.0.1',
  Port: 8000
)

server.mount_proc '/hello' do |req, res|
  res.body = 'Hello, World!'
end

server.mount_proc '/ping' do |req, res|
  res.body = 'pong!'
end

trap("INT"){ server.shutdown }

server.start
