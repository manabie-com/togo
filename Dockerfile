FROM ruby:3.0.2

RUN apt-get update -yqq \
  && apt-get install -yqq --no-install-recommends \
    postgresql-client \
    nodejs \
  && apt-get -q clean \
  && rm -rf /var/lib/apt/lists

  WORKDIR /usr/src/app
  COPY Gemfile* /usr/src/app/
  RUN bundle install

  COPY entrypoint.sh /usr/bin/
  RUN chmod +x /usr/bin/entrypoint.sh
  ENTRYPOINT ["entrypoint.sh"]
  EXPOSE 3000

  CMD ["rails", "server", "-b", "0.0.0.0"]