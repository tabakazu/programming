### Run chef-client.
### $ bundle exec knife zero converge "name:*" -o recipe[nginx::install] --why-run
### $ bundle exec knife zero converge "name:*" -o recipe[nginx::install]

case node["platform"]
when "amazon"
  if node["platform_version"] == "2"
    bash "install nginx" do
      code "amazon-linux-extras install nginx1.12"
      not_if { File.exists? "/usr/sbin/nginx" }
    end
  end
else
  package "nginx" do
    action :install
  end
end

service 'nginx' do
  action [:enable, :start]
end