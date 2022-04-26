module PostRequestTrackable
  extend ActiveSupport::Concern

  def count_request
    redis.set(client_id, 0, ex: 1.day) if redis.get(client_id).nil?
    redis.incr(client_id)
  end

  def client_reach_daily_limit?
    redis.get(client_id).to_i >= client_post_request_daily_limit
  end

  def current_user
    User.find_or_create_by(remote_ip: client_remote_ip) # todo: save to session
  end

  private

  def client_remote_ip
    request.remote_ip
  end

  def redis
    @redis ||= Redis.new
  end

  def client_id
    "#{client_remote_ip}:#{date}"
  end

  def date
    Time.now.strftime("%d:%m:%Y")
  end

  def client_post_request_daily_limit
    ENV["base_post_request_daily_limit"].to_i + current_user.id
  end
end
