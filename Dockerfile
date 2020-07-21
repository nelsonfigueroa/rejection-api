FROM ruby:2.7.1-alpine3.11

# ARG RUBYOPT='-W:no-deprecated -W:no-experimental'
# ENV RUBYOPT=$RUBYOPT
# ENV RAILS_ENV=production
# ENV RAILS_SERVE_STATIC_FILES=true
# ENV RAILS_LOG_TO_STDOUT=true

# to be able to run destructive database commands in production
# ENV DISABLE_DATABASE_ENVIRONMENT_CHECK=1

RUN apk update

# for nokogiri
RUN apk add build-base
# sqlite
RUN apk add sqlite-dev

WORKDIR /app

COPY Gemfile Gemfile.lock ./
RUN gem install bundler:2.1.4

# don't install test gems
RUN bundle config set without 'test'
RUN bundle install

RUN rm -rf /usr/local/bundle/cache/*.gem \
	&& find /usr/local/bundle/gems/ -name "*.c" -delete \
	&& find /usr/local/bundle/gems/ -name "*.o" -delete

COPY . .

EXPOSE 3000
CMD ["rails", "s", "-b", "0.0.0.0"]