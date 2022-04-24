module RequestTrackable
  extend ActiveSupport::Concern

  def count_request
    redis.set(client_id, 0, ex: 1.day) if redis.get(client_id).nil?
    redis.incr(client_id)
  end

  def client_reach_daily_limit?
    redis.get(client_id).to_i >= ENV["unauthenticated_post_todo_limit"].to_i
  end

  private

  def client_remote_ip
    request.remote_ip
  end

  def redis
    @redis ||= Redis.new
  end

  def client_id
    "#{request.remote_ip}:#{date}" # todo: change if user is authed
  end

  def date
    Time.now.strftime("%d:%m:%Y")
  end
end
