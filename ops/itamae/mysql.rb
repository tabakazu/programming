%w(
  mariadb
  mariadb-devel
  mariadb-libs
).each do |_package|
  package _package do
    action :install
  end
end

service 'mysql' do
  action [:enable, :start]
end