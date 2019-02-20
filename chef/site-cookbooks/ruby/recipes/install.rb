### Run chef-client.
### $ bundle exec knife zero converge "name:*" -o recipe[ruby::install] --why-run
### $ bundle exec knife zero converge "name:*" -o recipe[ruby::install]

%w(
  git 
  gcc
  gcc-c++
  openssl-devel 
  readline-devel 
  zlib-devel
).each do |_package|
  package _package do
    action :install
  end
end

git "/usr/local/rbenv" do
  repository "https://github.com/rbenv/rbenv.git"
  action :sync
end

directory "/usr/local/rbenv/plugins" do
  action :create
end

git "/usr/local/rbenv/plugins/ruby-build" do
  repository "https://github.com/rbenv/ruby-build.git"
  action :sync
end

template "/etc/profile.d/rbenv.sh" do
  source "/etc/profile.d/rbenv.sh.erb"
end

bash "install ruby" do
  code <<-EOS
    source /etc/profile.d/rbenv.sh;
    rbenv install -v 2.5.1
  EOS
  not_if { File.exists? "/usr/local/rbenv/versions/2.5.1" }
end

bash "set global ruby version" do
  code <<-EOS
    source /etc/profile.d/rbenv.sh;
    rbenv global 2.5.1;
    rbenv rehash
  EOS
  not_if <<-EOS
    source /etc/profile.d/rbenv.sh;
    ruby -v | grep 2.5.1
  EOS
end

bash "install bundler" do
  code <<-EOS
    source /etc/profile.d/rbenv.sh;
    gem install bundler -v '< 2.0.0'
  EOS
  not_if <<-EOS
    source /etc/profile.d/rbenv.sh;
    gem list | grep bundler
  EOS
end