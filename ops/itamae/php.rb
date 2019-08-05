%w(
  php
  php-cli
  php-devel
  php-mbstring
  php-mcrypt
  php-mysqlnd
  php-xml
).each do |_package|
  package _package do
    action :install
  end
end