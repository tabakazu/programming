%w(
  git
  htop
  curl
  wget
  nc
  vim
  unzip
  tcpdump
).each do |_package|
  package _package do
    action :install
  end
end
