require 'sinatra'

current_url_map = Hash.new

def gen_key(params)
  user = params[:user]
  topic = params[:topic]
  "%s@%s" % [user, topic]
end

put '/:user/:topic' do
  response.headers['Access-Control-Allow-Methods'] = '"Access-Control-Allow-Methods" ":" PUT'
  url = params[:url]
  key = gen_key(params)
  current_url_map[key] = url
end

get '/:user/:topic' do
  response.headers['Access-Control-Allow-Origin'] = '*'
  key = gen_key(params)
  url = current_url_map[key]
  url ||= "Not Found" # FIXME
end

# get '/:user/now' do
#   redirect last_topic, 303
# end
