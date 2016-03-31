require "sinatra"

configure { set :server, :puma }
puts "starting"

get "/" do
  "Hello cf-local-push!\n"
end
