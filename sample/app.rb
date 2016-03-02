require "sinatra"

configure { set :server, :puma }
puts "starting"

get "/" do
  "Hello cf-local-push!\n"
end

get "/env" do
  ENV.map do |k, v|
    "#{k} #{v}\n"
  end.join("")
end
