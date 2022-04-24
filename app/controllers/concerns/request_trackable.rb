module RequestTrackable
  extend ActiveSupport::Concern

  def count_request
    redis.incr(client_id)
  end

  def client_reach_daily_limit?
    redis.get(client_id).to_i >= 30 # todo: save as config
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
