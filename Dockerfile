FROM ruby:2.7.1-alpine3.11

RUN apk update
# for nokogiri
RUN apk add build-base
# sqlite
RUN apk add sqlite-dev

WORKDIR /app

COPY Gemfile Gemfile.lock ./
RUN gem install bundler:2.1.4

RUN bundle install

RUN rm -rf /usr/local/bundle/cache/*.gem \
	&& find /usr/local/bundle/gems/ -name "*.c" -delete \
	&& find /usr/local/bundle/gems/ -name "*.o" -delete

COPY . .

RUN rails db:setup

EXPOSE 3000
CMD ["rails", "s", "-b", "0.0.0.0"]