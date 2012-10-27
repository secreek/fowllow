require 'sinatra'

current_url_map = Hash.new

def gen_key(params)
  user = params[:user]
  meeting = params[:meeting]
  "%s@%s" % [user, meeting]
end

put '/:user/:meeting' do
  response.headers['Access-Control-Allow-Methods'] = '"Access-Control-Allow-Methods" ":" PUT'
  url = params[:url]
  key = gen_key(params)
  current_url_map[key] = url
end

get '/:user/:meeting' do
  response.headers['Access-Control-Allow-Origin'] = '*'
  key = gen_key(params)
  url = current_url_map[key]
  url ||= "Not Found"
end
