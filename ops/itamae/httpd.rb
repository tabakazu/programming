%w(
  httpd
  httpd-devel
  httpd-tools
).each do |_package|
  package _package do
    action :install
  end
end

service 'httpd' do
  action [:enable, :start]
end